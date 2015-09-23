package dao

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/lucasjo/porgex-go/db"
	"github.com/lucasjo/porgex-go/models"
)

func FindAllApps() []models.Application {

	conn := db.NewConnction()
	mSession := conn.Session()

	mSession.SetMode(mgo.Monotonic, true)
	defer mSession.Close()

	var apps []models.Application

	collection := mSession.DB(conn.Info().Database).C(C_application)

	err1 := collection.Find(nil).All(&apps)
	if err1 != nil {
		log.Printf("RunQuery : Error : %s\n", err1)
		return apps
	}

	return apps
}

func FindIdApp(id string) models.Application {
	conn := db.NewConnction()
	mSession := conn.Session()

	mSession.SetMode(mgo.Monotonic, true)
	defer mSession.Close()

	var app models.Application

	collection := mSession.DB(conn.Info().Database).C(C_application)
	oid := bson.ObjectIdHex(id)
	err1 := collection.FindId(oid).One(&app)
	if err1 != nil {
		log.Printf("RunQuery : Error : %s\n", err1)
		return app
	}

	return app
}
