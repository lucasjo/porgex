package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

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

	fmt.Printf("fields %v : ", fields)

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

func main() {
	id := "55ee3a460f5106ab680000ca"

	appCgroupPath := filepath.Join(appCpuAcctPath, id)

	user, system, _ := GetCpuUsageStat(appCgroupPath)

	fmt.Printf("user %v , system %v", user, system)

}
