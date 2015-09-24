package main

import (
	"fmt"

	"github.com/lucasjo/porgex/porgex-agent/models"
	"github.com/lucasjo/porgex/porgex-agent/usage/mem"
)

func main() {
	id := "55ee3a460f5106ab680000ca"

	var cStats = &models.AppStats{}

	err := usage.SetMemoryStats(id, cStats)

	if err != nil {
		fmt.Errorf("Error : ", err)
	}

	fmt.Printf("stats : %v\n", cStats)
}
