package main

import (
	"fmt"

	"github.com/lucasjo/porgex/go-collect-server/config"
	//"github.com/olebedev/config"
)

const ConfigName = "collect_server.yaml"

func main() {
	cfg := config.GetConfig("")
	port, _ := cfg.String("development.tcp.port")
	fmt.Printf("config %v\n", port)
}
