package usage

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/lucasjo/porgex/porgex-agent/system"
)

var appCpuAcctPath = "/cgroup/cpuacct/openshift"

const (
	cpuacctStat        = "cpuacct.stat"
	nanosecondInSecond = 1000000000
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

	data, err := ioutil.RedaFile(filepath.Join(path, cpuacctStat))

	if err != nil {
		return 0, 0, err
	}
	fields := strings.Fields(string(data))
	fmt.PrintF("fields : %v ", fields)

	if len(fields) != 4 {
		return 0, 0, fmt.Errorf("failure - %s is expected to have 4 fields", filepath.Join(path, cpuacctStat))
	}

}
