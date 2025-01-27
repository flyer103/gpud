// Package process provides the process runner implementation on the host.
package process

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/leptonai/gpud/log"
)

type Process interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error

	// Waits for the process to exit and returns the error, if any.
	// If the command completes successfully, the error will be nil.
	Wait() <-chan error

	PID() int32

	StdoutReader() io.Reader
	StderrReader() io.Reader
}

// RestartConfig is the configuration for the process restart.
type RestartConfig struct {
	// Set true to restart the process on error exit.
	OnError bool
	// Set the maximum number of restarts.
	Limit int
	// Set the interval between restarts.
	Interval time.Duration
}

type process struct {
	ctx    context.Context
	cancel context.CancelFunc

	cmdMu       sync.RWMutex
	cmd         *exec.Cmd
	errc        chan error
	pid         int32
	commandArgs []string
	envs        []string
	runBashFile *os.File

	outputFile   *os.File
	stdoutReader io.ReadCloser
	stderrReader io.ReadCloser

	wg sync.WaitGroup

	restartConfig *RestartConfig
}

func New(commands [][]string, opts ...OpOption) (Process, error) {
	op := &Op{}
	if err := op.applyOpts(opts); err != nil {
		return nil, err
	}
	if len(commands) == 0 {
		return nil, errors.New("no commands provided")
	}
	if len(commands) > 1 && !op.runAsBashScript {
		return nil, errors.New("cannot run multiple commands without a bash script mode")
	}
	for _, args := range commands {
		cmd := strings.Split(args[0], " ")[0]
		if !commandExists(cmd) {
			return nil, fmt.Errorf("command not found: %q", cmd)
		}
	}

	var cmdArgs []string
	var bashFile *os.File
	if op.runAsBashScript {
		var err error
		bashFile, err = os.CreateTemp(os.TempDir(), "tmpbash*.bash")
		if err != nil {
			return nil, err
		}
		if _, err := bashFile.Write([]byte(bashScriptHeader)); err != nil {
			return nil, err
		}
		defer func() {
			_ = bashFile.Sync()
		}()
		cmdArgs = []string{"bash", bashFile.Name()}
	}

	for _, args := range commands {
		if bashFile == nil {
			cmdArgs = args
			continue
		}

		if _, err := bashFile.Write([]byte(strings.Join(args, " "))); err != nil {
			return nil, err
		}
		if _, err := bashFile.Write([]byte("\n")); err != nil {
			return nil, err
		}
	}

	errcBuffer := 1
	if op.restartConfig != nil && op.restartConfig.OnError && op.restartConfig.Limit > 0 {
		errcBuffer = op.restartConfig.Limit
	}
	return &process{
		cmd:         nil,
		errc:        make(chan error, errcBuffer),
		commandArgs: cmdArgs,
		envs:        op.envs,
		runBashFile: bashFile,
		outputFile:  op.outputFile,

		restartConfig: op.restartConfig,
	}, nil
}

func (p *process) Start(ctx context.Context) error {
	p.cmdMu.Lock()
	defer p.cmdMu.Unlock()

	if p.cmd != nil {
		return errors.New("process already started")
	}

	cctx, ccancel := context.WithCancel(ctx)
	p.ctx = cctx
	p.cancel = ccancel

	if err := p.startCommand(); err != nil {
		return err
	}

	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		p.cmdWait()
	}()

	return nil
}

func (p *process) startCommand() error {
	log.Logger.Debugw("starting command", "command", p.commandArgs)
	p.cmd = exec.CommandContext(p.ctx, p.commandArgs[0], p.commandArgs[1:]...)
	p.cmd.Env = p.envs

	switch {
	case p.outputFile != nil:
		p.cmd.Stdout = p.outputFile
		p.cmd.Stderr = p.outputFile

	default:
		var err error
		p.stdoutReader, err = p.cmd.StdoutPipe()
		if err != nil {
			return fmt.Errorf("failed to get stdout pipe: %w", err)
		}
		p.stderrReader, err = p.cmd.StderrPipe()
		if err != nil {
			return fmt.Errorf("failed to get stderr pipe: %w", err)
		}
	}

	if err := p.cmd.Start(); err != nil {
		return fmt.Errorf("failed to start command: %w", err)
	}
	atomic.StoreInt32(&p.pid, int32(p.cmd.Process.Pid))

	return nil
}

func (p *process) Wait() <-chan error {
	return p.errc
}

func (p *process) cmdWait() {
	restartCount := 0
	for {
		errc := make(chan error)
		go func() {
			errc <- p.cmd.Wait()
		}()

		select {
		case <-p.ctx.Done():
			// command aborted (e.g., Stop called)
			// cmd.Wait will return error
			err := <-errc
			p.errc <- err
			return

		case err := <-errc:
			p.errc <- err

			if err == nil {
				log.Logger.Debugw("process exited successfully")
				return
			}

			if exitErr, ok := err.(*exec.ExitError); ok {
				if exitErr.ExitCode() == -1 {
					if p.ctx.Err() != nil {
						log.Logger.Debugw("command was terminated (exit code -1) by the root context cancellation", "cmd", p.cmd.String(), "contextError", p.ctx.Err())
					} else {
						log.Logger.Warnw("command was terminated (exit code -1) for unknown reasons", "cmd", p.cmd.String())
					}
				} else {
					log.Logger.Warnw("command exited with non-zero status", "error", err, "cmd", p.cmd.String(), "exitCode", exitErr.ExitCode())
				}
			} else {
				log.Logger.Warnw("error waiting for command to finish", "error", err, "cmd", p.cmd.String())
			}

			if p.restartConfig == nil || !p.restartConfig.OnError {
				log.Logger.Warnw("process exited with error", "error", err)
				return
			}

			if p.restartConfig.Limit > 0 && restartCount >= p.restartConfig.Limit {
				log.Logger.Warnw("process exited with error, but restart limits reached", "restartCount", restartCount, "error", err)
				return
			}
		}

		select {
		case <-p.ctx.Done():
			return
		case <-time.After(p.restartConfig.Interval):
		}

		if err := p.startCommand(); err != nil {
			log.Logger.Warnw("failed to restart command", "error", err)
			return
		}

		restartCount++
	}
}

func (p *process) Stop(ctx context.Context) error {
	p.cmdMu.Lock()
	defer p.cmdMu.Unlock()

	if p.cmd == nil {
		return errors.New("process not started")
	}

	p.cancel()

	finished := false
	if err := p.cmd.Process.Signal(syscall.SIGTERM); err != nil {
		if err.Error() == "os: process already finished" {
			finished = true
		} else {
			log.Logger.Warnw("failed to send SIGTERM to process", "error", err)
		}
	}

	if !finished {
		select {
		case <-p.ctx.Done():
			return ctx.Err()
		case <-time.After(3 * time.Second):
			if err := p.cmd.Process.Kill(); err != nil {
				log.Logger.Warnw("failed to send SIGKILL to process", "error", err)
			}
		}
	}

	if p.runBashFile != nil {
		_ = p.runBashFile.Sync()
		_ = p.runBashFile.Close()
		return os.RemoveAll(p.runBashFile.Name())
	}

	p.cmd = nil
	return nil
}

func (p *process) PID() int32 {
	return atomic.LoadInt32(&p.pid)
}

func (p *process) StdoutReader() io.Reader {
	p.cmdMu.RLock()
	defer p.cmdMu.RUnlock()

	if p.outputFile != nil {
		return p.outputFile
	}
	return p.stdoutReader
}

func (p *process) StderrReader() io.Reader {
	p.cmdMu.RLock()
	defer p.cmdMu.RUnlock()

	if p.outputFile != nil {
		return p.outputFile
	}
	return p.stderrReader
}

const bashScriptHeader = `#!/bin/bash

# do not mask errors in a pipeline
set -o pipefail

# treat unset variables as an error
set -o nounset

# exit script whenever it errs
set -o errexit

`

func commandExists(name string) bool {
	p, err := exec.LookPath(name)
	if err != nil {
		return false
	}
	return p != ""
}
