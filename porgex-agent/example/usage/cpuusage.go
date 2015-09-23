package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/lucasjo/porgex/porgex-agent/models"
	"github.com/lucasjo/porgex/porgex-agent/system"
)

var appCpuAcctPath = "/cgroup/cpuacct/openshift/"

const (
	cpuacctStat         = "cpuacct.stat"
	nanosecondsInSecond = 1000000000
)

var clockTicks = uint64(system.GetClockTicks())

func GetCpuUsageStat(path string) (uint64, uint64, error) {
	userModeUsage := uint64(0)
	systemModeUsage := uint64(0)

	const (
		userField   = "user"
		systemField = "system"
	)

	// user <usage in ticks>
	// system <usage in ticks>

	data, err := ioutil.ReadFile(filepath.Join(path, cpuacctStat))

	if err != nil {
		return 0, 0, err
	}
	fields := strings.Fields(string(data))

	fmt.Printf("fields : %v\n", fields)

	if len(fields) != 4 {
		return 0, 0, fmt.Errorf("failure - %s is expected to have 4 fields", filepath.Join(path, cpuacctStat))
	}

	if fields[0] != userField {
		return 0, 0, fmt.Errorf("unexpected field %q in %q, expected %q", fields[0], cpuacctStat, userField)
	}
	if fields[2] != systemField {
		return 0, 0, fmt.Errorf("unexpected field %q in %q, expected %q", fields[2], cpuacctStat, systemField)
	}
	if userModeUsage, err = strconv.ParseUint(fields[1], 10, 64); err != nil {
		return 0, 0, err
	}

	if systemModeUsage, err = strconv.ParseUint(fields[3], 10, 64); err != nil {
		return 0, 0, err
	}

	return (userModeUsage * nanosecondsInSecond) / clockTicks, (systemModeUsage * nanosecondsInSecond) / clockTicks, nil

}

func SetCpuUsage(uuid string, stats *models.AppCpuStats) error {

	appCgroupPath := filepath.Join(appCpuAcctPath, uuid)

	userModeUsage, systemModeUsage, err := GetCpuUsageStat(appCgroupPath)

	if err != nil {
		return err
	}

	totalUsage, err := getUsageUint(appCgroupPath, "cpuacct.usage")

	if err != nil {
		fmt.Errorf("Error : ", err)
		return err
	}

	perCpuUsage, err := getPercpuUsage(appCgroupPath)

	if err != nil {
		return err
	}

	stats.CPUStats.CPUUsage.TotalUsage = totalUsage
	stats.CPUStats.CPUUsage.UsageInSytemmode = systemModeUsage
	stats.CPUStats.CPUUsage.UsageInUsermode = userModeUsage
	stats.CPUStats.CPUUsage.PercpuUsage = perCpuUsage

	return nil
}

func getUsageUint(path, param string) (uint64, error) {

	contents, err := ioutil.ReadFile(filepath.Join(path, param))

	if err != nil {
		return 0, err
	}

	return ParseUint(strings.TrimSpace(string(contents)), 10, 64)

}

func getPercpuUsage(path string) ([]uint64, error) {
	percpuUsage := []uint64{}
	data, err := ioutil.ReadFile(filepath.Join(path, "cpuacct.usage_percpu"))
	if err != nil {
		return percpuUsage, err
	}
	for _, value := range strings.Fields(string(data)) {
		value, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return percpuUsage, fmt.Errorf("Unable to convert param value to uint64: %s", err)
		}
		percpuUsage = append(percpuUsage, value)
	}
	return percpuUsage, nil
}

func ParseUint(s string, base, bitSize int) (uint64, error) {
	value, err := strconv.ParseUint(s, base, bitSize)
	if err != nil {
		intValue, intErr := strconv.ParseInt(s, base, bitSize)

		if intErr == nil && intValue < 0 {
			return 0, nil

		} else if intErr != nil && intErr.(*strconv.NumError).Err == strconv.ErrRange && intValue < 0 {
			return 0, nil
		}
		return value, err
	}
	return value, err
}

func calculateCPUPercent(previousCpu uint64, v *models.AppCpuStats) float64 {

	var cpuPercnt = 0.0

	cpuPercnt = (float64(previousCpu) / float64(v.CPUStats.CPUUsage.TotalUsage)) / float64(len(v.CPUStats.CPUUsage.PercpuUsage)) * 100.0

	return cpuPercnt

}

func main() {
	id := "55ee3a460f5106ab680000ca"

	var cStats = &models.AppCpuStats{}

	err := SetCpuUsage(id, cStats)

	if err != nil {
		fmt.Errorf("error message : ", err)
	}

	var previousCpu = cStats.CPUStats.CPUUsage.TotalUsage

	time.Sleep(time.Second * 2)

	err1 := SetCpuUsage(id, cStats)

	if err1 != nil {
		fmt.Errorf("error message : ", err1)
	}

	cpuPer := calculateCPUPercent(previousCpu, cStats)

	fmt.Printf("stats %v\n", cpuPer)

}
