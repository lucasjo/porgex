package main

import (
	"os/user"

	log "github.com/Sirupsen/logrus"
	"github.com/lucasjo/porgex/go-collect-server/config"
	"github.com/lucasjo/porgex/go-collect-server/logger"
	"github.com/lucasjo/porgex/go-collect-server/server"
)

func main() {
	configpath := "/Users/kikimans/go/src/github.com/lucasjo/porgex/go-collect-server/collect_config.yaml"
	cfg := config.GetConfig(configpath)

	loglevel, _ := cfg.String("development.common.loglevel")

	if err := logger.SetLevel(loglevel); err != nil {
		log.Fatal(err)
	}

	logfile, _ := cfg.String("development.common.logfile")

	var runAsUser *user.User

	if err := logger.PrepareFile(logfile, runAsUser); err != nil {
		log.Fatal(err)
	}

	if err := logger.SetFile(logfile); err != nil {
		log.Fatal(err)
	}

	/*
		runtime.LockOSThread()

		cntxt := &daemon.Context{
			PidFileName: "collect-server.pid",
			PidFilePerm: 0644,
			WorkDir:     "./",
			Umask:       027,
			Args:        []string{"[porgex-agent] collect "},
		}

		child, _ := cntxt.Reborn()

		if child != nil {
			return

		}
		defer cntxt.Release()
		runtime.UnlockOSThread()
	*/
	s := server.NewServer()
	defer s.Stop()

	s.Listen()

}
