// Package ecc implements NVIDIA GPU ECC error monitoring.
package ecc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/leptonai/gpud/components"
	nvidia_query "github.com/leptonai/gpud/components/accelerator/nvidia/query"
	nvidia_query_metrics_ecc "github.com/leptonai/gpud/components/accelerator/nvidia/query/metrics/ecc"
	"github.com/leptonai/gpud/components/query"
	"github.com/leptonai/gpud/log"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	Name        = "accelerator-nvidia-ecc"
	Description = "Tracks the per-GPU ECC errors."
)

var Tags = []string{"nvidia", "gpu", "ecc"}

func New(ctx context.Context, cfg Config) components.Component {
	cfg.Query.SetDefaultsIfNotSet()

	cctx, ccancel := context.WithCancel(ctx)
	nvidia_query.DefaultPoller.Start(cctx, cfg.Query, Name)

	return &component{
		rootCtx: ctx,
		cancel:  ccancel,
		poller:  nvidia_query.DefaultPoller,
	}
}

var _ components.Component = (*component)(nil)

type component struct {
	rootCtx  context.Context
	cancel   context.CancelFunc
	poller   query.Poller
	gatherer prometheus.Gatherer
}

func (c *component) Name() string { return Name }

func (c *component) States(ctx context.Context) ([]components.State, error) {
	last, err := c.poller.Last()
	if err != nil {
		return nil, err
	}
	if last == nil { // no data
		log.Logger.Debugw("nothing found in last state (no data collected yet)", "component", Name)
		return nil, nil
	}
	if last.Error != nil {
		return []components.State{
			{
				Healthy: false,
				Error:   last.Error,
				Reason:  "last query failed",
			},
		}, nil
	}
	if last.Output == nil {
		return []components.State{
			{
				Healthy: false,
				Reason:  "no output",
			},
		}, nil
	}

	allOutput, ok := last.Output.(*nvidia_query.Output)
	if !ok {
		return nil, fmt.Errorf("invalid output type: %T", last.Output)
	}
	if allOutput.SMIExists && len(allOutput.SMIQueryErrors) > 0 {
		cs := make([]components.State, 0)
		for _, e := range allOutput.SMIQueryErrors {
			cs = append(cs, components.State{
				Name:    Name,
				Healthy: false,
				Error:   errors.New(e),
				Reason:  "nvidia-smi query failed with " + e,
				ExtraInfo: map[string]string{
					nvidia_query.StateKeySMIExists: fmt.Sprintf("%v", allOutput.SMIExists),
				},
			})
		}
		return cs, nil
	}
	output := ToOutput(allOutput)
	return output.States()
}

func (c *component) Events(ctx context.Context, since time.Time) ([]components.Event, error) {
	return nil, nil
}

func (c *component) Metrics(ctx context.Context, since time.Time) ([]components.Metric, error) {
	log.Logger.Debugw("querying metrics", "since", since)

	aggTotalCorrecteds, err := nvidia_query_metrics_ecc.ReadAggregateTotalCorrected(ctx, since)
	if err != nil {
		return nil, fmt.Errorf("failed to read aggregate total corrected: %w", err)
	}
	aggTotalUncorrecteds, err := nvidia_query_metrics_ecc.ReadAggregateTotalUncorrected(ctx, since)
	if err != nil {
		return nil, fmt.Errorf("failed to read aggregate total corrected: %w", err)
	}
	volTotalCorrecteds, err := nvidia_query_metrics_ecc.ReadVolatileTotalCorrected(ctx, since)
	if err != nil {
		return nil, fmt.Errorf("failed to read volatile total corrected: %w", err)
	}
	volTotalUncorrecteds, err := nvidia_query_metrics_ecc.ReadVolatileTotalUncorrected(ctx, since)
	if err != nil {
		return nil, fmt.Errorf("failed to read volatile total corrected: %w", err)
	}

	ms := make([]components.Metric, 0, len(aggTotalCorrecteds)+len(aggTotalUncorrecteds)+len(volTotalCorrecteds)+len(volTotalUncorrecteds))
	for _, m := range aggTotalCorrecteds {
		ms = append(ms, components.Metric{
			Metric: m,
			ExtraInfo: map[string]string{
				"gpu_id": m.MetricSecondaryName,
			},
		})
	}
	for _, m := range aggTotalUncorrecteds {
		ms = append(ms, components.Metric{
			Metric: m,
			ExtraInfo: map[string]string{
				"gpu_id": m.MetricSecondaryName,
			},
		})
	}
	for _, m := range volTotalCorrecteds {
		ms = append(ms, components.Metric{
			Metric: m,
			ExtraInfo: map[string]string{
				"gpu_id": m.MetricSecondaryName,
			},
		})
	}
	for _, m := range volTotalUncorrecteds {
		ms = append(ms, components.Metric{
			Metric: m,
			ExtraInfo: map[string]string{
				"gpu_id": m.MetricSecondaryName,
			},
		})
	}

	return ms, nil
}

func (c *component) Close() error {
	log.Logger.Debugw("closing component")

	// safe to call stop multiple times
	_ = c.poller.Stop(Name)

	return nil
}

var _ components.PromRegisterer = (*component)(nil)

func (c *component) RegisterCollectors(reg *prometheus.Registry, db *sql.DB, tableName string) error {
	c.gatherer = reg
	return nvidia_query_metrics_ecc.Register(reg, db, tableName)
}