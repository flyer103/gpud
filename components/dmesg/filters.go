package dmesg

import (
	"github.com/leptonai/gpud/components/memory"
	query_log_filter "github.com/leptonai/gpud/components/query/log/filter"

	"k8s.io/utils/ptr"
)

const (
	// e.g.,
	// Out of memory: Killed process 123, UID 48, (httpd).
	EventOOMKill      = "oom_kill"
	EventOOMKillRegex = `Out of memory:`

	// e.g.,
	// postgres invoked oom-killer: gfp_mask=0x201d2, order=0, oomkilladj=0
	EventOOMKiller      = "oom_killer"
	EventOOMKillerRegex = `(?i)\b(invoked|triggered) oom-killer\b`

	// e.g.,
	// Memory cgroup out of memory: Killed process 123, UID 48, (httpd).
	EventOOMCgroup      = "oom_cgroup"
	EventOOMCgroupRegex = `Memory cgroup out of memory`
)

var defaultFilters = []*query_log_filter.Filter{
	{
		Name:            EventOOMKill,
		Regex:           ptr.To(EventOOMKillRegex),
		OwnerReferences: []string{memory.Name},
	},
	{
		Name:            EventOOMKiller,
		Regex:           ptr.To(EventOOMKillerRegex),
		OwnerReferences: []string{memory.Name},
	},
	{
		Name:            EventOOMCgroup,
		Regex:           ptr.To(EventOOMCgroupRegex),
		OwnerReferences: []string{memory.Name},
	},
}

func DefaultLogFilters() []*query_log_filter.Filter {
	return defaultFilters
}
