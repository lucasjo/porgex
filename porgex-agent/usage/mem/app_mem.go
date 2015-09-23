package mem

import (
	"fmt"
	"os"
	"strings"
)

type MemoryData struct {
	MaxUsage uint64
	Limit    uint64
	Usage    uint64
}

var openshiftCgroupsMemName = "/cgroup/memory/openshift"

func GetAppMemory(uuid string) (MemoryData, error) {

	var mData MemoryData

	moduleName := "memory"

	openshiftCgroupsMemName = strings.Join([]string{openshiftCgroupsMemName, uuid}, "/")

	usage := strings.Join([]string{moduleName, "usage_in_bytes"}, ".")
	maxUsage := strings.Join([]string{moduleName, "max_usage_in_bytes"}, ".")
	limit := strings.Join([]string{moduleName, "limit_in_bytes"}, ".")

	usageValue, err := getCgroupModuleInfo(openshiftCgroupsMemName, usage)

	if err != nil {
		if moduleName != "memory" && os.IsNotExist(err) {
			return MemoryData{}, nil
		}
		return MemoryData{}, fmt.Errorf("failed to parse %s - %v\n", usage, err)
	}

	mData.Usage = usageValue

	maxUusageValue, err := getCgroupModuleInfo(openshiftCgroupsMemName, maxUsage)

	if err != nil {
		if moduleName != "memory" && os.IsNotExist(err) {
			return MemoryData{}, nil
		}
		return MemoryData{}, fmt.Errorf("failed to parse %s - %v\n", maxUusageValue, err)
	}

	mData.MaxUsage = maxUusageValue

	limitValue, err := getCgroupModuleInfo(openshiftCgroupsMemName, limit)

	if err != nil {
		if moduleName != "memory" && os.IsNotExist(err) {
			return MemoryData{}, nil
		}
		return MemoryData{}, fmt.Errorf("failed to parse %s - %v\n", limit, err)
	}

	mData.Limit = limitValue

	return mData, nil

}
