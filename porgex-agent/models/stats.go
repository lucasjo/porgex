package models

// CPUUsage stores All CPU stats aggregated since container inception.
type CPUUsage struct {
	// Total CPU time consumed.
	// Units: nanoseconds.
	TotalUsage uint64 `json:"total_usage"`
	// Total CPU time consumed per core.
	// Units: nanoseconds.
	PercpuUsage []uint64 `json:"percpu_usage"`
	// Time spent by tasks of the cgroup in kernel mode.
	// Units: nanoseconds.
	UsageInSytemmode uint64 `json:"usage_in_kernelmode"`
	// Time spent by tasks of the cgroup in user mode.
	// Units: nanoseconds.
	UsageInUsermode uint64 `json:"usage_in_usermode"`
}

// CPUStats aggregates and wraps all CPU related info of container
type CPUStats struct {
	CPUUsage    CPUUsage `json:"cpu_usage"`
	SystemUsage uint64   `json:"system_cpu_usage"`
}

type AppCpuStats struct {
	App      Application `json:"app"`
	CPUStats CPUStats    `json:"cpu_stats,omitempty"`
}

type MemStats struct {
	MaxUsage     uint64 `json:"max_usage"`
	LimitUsage   uint64 `json:"limit_usage"`
	CurrentUsage uint64 `json:"current_usage"`
}

type AppStats struct {
	App      Application `json:"app"`
	CPUStats CPUStats    `json:"cpu_stats,omitempty"`
	MemStats MemStats    `json:"mem_stats, omitempty"`
}
