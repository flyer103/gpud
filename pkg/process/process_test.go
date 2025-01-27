package process

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestProcess(t *testing.T) {
	t.Parallel()

	p, err := New(
		[][]string{
			{"echo", "hello"},
		},
		WithOutputFile(os.Stderr),
	)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := p.Start(ctx); err != nil {
		t.Fatal(err)
	}
	t.Logf("pid: %d", p.PID())

	select {
	case err := <-p.Wait():
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("timeout")
	}

	if err := p.Stop(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestProcessWithBash(t *testing.T) {
	t.Parallel()

	p, err := New(
		[][]string{
			{"echo", "hello"},
			{"echo hello && echo 111 | grep 1"},
		},
		WithOutputFile(os.Stderr),
		WithRunAsBashScript(),
	)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := p.Start(ctx); err != nil {
		t.Fatal(err)
	}
	t.Logf("pid: %d", p.PID())

	select {
	case err := <-p.Wait():
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("timeout")
	}

	if err := p.Stop(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestProcessWithTempFile(t *testing.T) {
	t.Parallel()

	// create a temporary file
	tmpFile, err := os.CreateTemp("", "process-test-*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	p, err := New(
		[][]string{
			{"echo", "hello"},
		},
		WithOutputFile(tmpFile),
	)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := p.Start(ctx); err != nil {
		t.Fatal(err)
	}
	t.Logf("pid: %d", p.PID())

	select {
	case err := <-p.Wait():
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("timeout")
	}

	if err := p.Stop(ctx); err != nil {
		t.Fatal(err)
	}

	// Verify the content of the temporary file
	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	expectedContent := "hello\n"
	if string(content) != expectedContent {
		t.Fatalf("Expected content %q, but got %q", expectedContent, string(content))
	}
}

func TestProcessWithStdoutReader(t *testing.T) {
	t.Parallel()

	p, err := New(
		[][]string{
			{"echo hello && sleep 1000"},
		},
		WithRunAsBashScript(),
	)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := p.Start(ctx); err != nil {
		t.Fatal(err)
	}
	t.Logf("pid: %d", p.PID())

	select {
	case err := <-p.Wait():
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(time.Second):
	}

	rd := p.StdoutReader()
	buf := make([]byte, 1024)
	n, err := rd.Read(buf)
	if err != nil {
		t.Fatal(err)
	}
	output := string(buf[:n])
	expectedOutput := "hello\n"
	if output != expectedOutput {
		t.Fatalf("expected output %q, but got %q", expectedOutput, output)
	}
	t.Logf("stdout: %q", output)

	if err := p.Stop(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestProcessWithStdoutReaderUntilEOF(t *testing.T) {
	t.Parallel()

	p, err := New(
		[][]string{
			{"echo hello 1 && sleep 1"},
			{"echo hello 2 && sleep 1"},
			{"echo hello 3 && sleep 1"},
		},
		WithRunAsBashScript(),
	)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := p.Start(ctx); err != nil {
		t.Fatal(err)
	}
	t.Logf("pid: %d", p.PID())

	rd := p.StdoutReader()
	scanner := bufio.NewScanner(rd)
	var output string
	for scanner.Scan() {
		output += scanner.Text() + "\n"
	}
	expectedOutput := "hello 1\nhello 2\nhello 3\n"
	if output != expectedOutput {
		t.Fatalf("expected output %q, but got %q", expectedOutput, output)
	}
	t.Logf("stdout: %q", output)

	select {
	case err := <-p.Wait():
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(time.Second):
	}

	if err := p.Stop(ctx); err != nil {
		t.Fatal(err)
	}
	if scanner.Err() != nil {
		t.Fatal(scanner.Err())
	}
}

func TestProcessWithRestarts(t *testing.T) {
	t.Parallel()

	p, err := New(
		[][]string{
			{"echo hello"},
			{"echo 111 && exit 1"},
		},
		WithOutputFile(os.Stderr),
		WithRunAsBashScript(),
		WithRestartConfig(RestartConfig{
			OnError:  true,
			Limit:    3,
			Interval: 100 * time.Millisecond,
		}),
	)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := p.Start(ctx); err != nil {
		t.Fatal(err)
	}
	t.Logf("pid: %d", p.PID())

	for i := 0; i < 3; i++ {
		select {
		case err := <-p.Wait():
			if err == nil {
				t.Fatal("expected error")
			}
			if strings.Contains(err.Error(), "exit status 1") {
				t.Log(err)
				continue
			}
			t.Fatal(err)

		case <-time.After(2 * time.Second):
			t.Fatal("timeout")
		}
	}

	if err := p.Stop(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestProcessSleep(t *testing.T) {
	t.Parallel()

	p, err := New(
		[][]string{
			{"sleep", "9999"},
		},
		WithOutputFile(os.Stderr),
	)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := p.Start(ctx); err != nil {
		t.Fatal(err)
	}
	t.Logf("pid: %d", p.PID())

	if err := p.Stop(ctx); err != nil {
		t.Fatal(err)
	}

	select {
	case err := <-p.Wait():
		if err == nil {
			t.Fatal("expected error")
		}
		t.Log(err)
	case <-time.After(3 * time.Second):
		t.Fatal("timeout")
	}
}

func TestProcessStream(t *testing.T) {
	t.Parallel()

	cmds := make([][]string, 0, 100)
	for i := 0; i < 100; i++ {
		cmds = append(cmds, []string{fmt.Sprintf("echo hello %d && sleep 1", i)})
	}

	p, err := New(cmds, WithRunAsBashScript())
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := p.Start(ctx); err != nil {
		t.Fatal(err)
	}
	t.Logf("pid: %d", p.PID())

	rd := p.StdoutReader()
	buf := make([]byte, 1024)
	for i := 0; i < 3; i++ {
		n, err := rd.Read(buf)
		if err != nil {
			t.Fatal(err)
		}

		output := string(buf[:n])
		expectedOutput := fmt.Sprintf("hello %d\n", i)
		if output != expectedOutput {
			t.Fatalf("expected output %q, but got %q", expectedOutput, output)
		}
		t.Logf("stdout: %q", output)
	}

	if err := p.Stop(ctx); err != nil {
		t.Fatal(err)
	}
}
