package main

import (
	"fmt"

	"github.com/lucasjo/porgex/porgex-agent/system"
)

func main() {
	fmt.Printf("system clock %d", system.GetClockTicks())
}
