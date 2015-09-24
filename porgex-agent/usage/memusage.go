package usage

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/lucasjo/porgex/porgex-agent/models"
)

var AppMemCgroupPath = "/cgroup/memory/openshift/"

const (
	memstat  = "memory"
	megaByte = 1024 * 1024
)

func SetMemoryStats(uuid string, v *models.AppStats) error {

	cgroupPath := filepath.Join(AppMemCgroupPath, uuid)

	usagefile := strings.Join([]string{memstat, "usage_in_bytes"}, ".")
	maxUsagefile := strings.Join([]string{memstat, "max_usage_in_bytes"}, ".")
	limitfile := strings.Join([]string{memstat, "limit_in_bytes"}, ".")

	usageValue, err := getUsageUint(cgroupPath, usagefile)

	if err != nil {
		fmt.Errorf("failed to parse %s - %v\n", usagefile, err)
		return err
	}

	maxUsageValue, err := getUsageUint(cgroupPath, maxUsagefile)

	if err != nil {
		fmt.Errorf("failed to parse %s - %v\n", maxUsagefile, err)
		return err
	}

	limitValue, err := getUsageUint(cgroupPath, limitfile)

	if err != nil {
		fmt.Errorf("failed to parse %s - %v\n", limitfile, err)
		return err
	}

	v.MemStats.CurrentUsage = usageValue
	v.MemStats.MaxUsage = maxUsageValue
	v.MemStats.LimitUsage = limitValue

	return nil

}
