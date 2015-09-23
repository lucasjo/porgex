package dao

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/lucasjo/porgex-go/db"
	"github.com/lucasjo/porgex-go/models"
)

func GetAllCloudUsers() []models.CloudUsers {

	conn := db.NewConnction()
	mSession := conn.Session()

	mSession.SetMode(mgo.Monotonic, true)
	defer mSession.Close()

	var cloudusers []models.CloudUsers

	collection := mSession.DB(conn.Info().Database).C(C_cloud_users)

	err1 := collection.Find(nil).All(&cloudusers)
	if err1 != nil {
		log.Printf("RunQuery : Error : %s\n", err1)
		return cloudusers
	}

	return cloudusers
}

func FindLoiginCloudUsers(login string) models.CloudUsers {

	conn := db.NewConnction()
	mSession := conn.Session()

	mSession.SetMode(mgo.Monotonic, true)
	defer mSession.Close()

	var clouduser models.CloudUsers

	collection := mSession.DB(conn.Info().Database).C(C_cloud_users)

	err1 := collection.Find(bson.M{"login": login}).One(&clouduser)
	if err1 != nil {
		log.Printf("RunQuery : Error : %s\n", err1)
		return clouduser
	}

	return clouduser

}

func FindIdCloudUsers(id string) models.CloudUsers {
	conn := db.NewConnction()
	mSession := conn.Session()

	mSession.SetMode(mgo.Monotonic, true)
	defer mSession.Close()

	var clouduser models.CloudUsers

	collection := mSession.DB(conn.Info().Database).C(C_cloud_users)
	oid := bson.ObjectIdHex(id)
	err1 := collection.FindId(oid).One(&clouduser)
	if err1 != nil {
		log.Printf("RunQuery : Error : %s\n", err1)
		return clouduser
	}

	return clouduser
}
