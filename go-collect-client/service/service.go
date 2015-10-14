package service

import (
	"fmt"
	"os"

	"github.com/lucasjo/porgex/go-collect-client/db"
	"github.com/lucasjo/porgex/porgex-agent/models"
)

func GetServerApplication() []models.Application {

	var apps []models.Application

	db.FindAll(db.Application, nil, &apps)

	hostname, err := os.Hostname()

	fmt.Println("hostname : ", hostname)

	if err != nil {
		fmt.Errorf("Error hostname get", err)
	}

	var rApps []models.Application
	for _, app := range apps {
		for _, gear := range app.Gears {
			if gear.Server_identity == hostname {
				rApps = append(rApps, app)
			}
		}
	}

	return rApps

}
