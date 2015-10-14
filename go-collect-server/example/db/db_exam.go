package main

import (
	"fmt"

	"github.com/lucasjo/porgex/go-collect-server/config"
	"github.com/lucasjo/porgex/go-collect-server/db"
)

func main() {
	c := db.New(config.GetConfig("/Users/kikimans/go/src/github.com/lucasjo/porgex/go-collect-server/collect_config.yaml"))
	fmt.Printf("connect %v\n", c)

}
