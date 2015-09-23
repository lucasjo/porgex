package db

import (
	"os"
	"path/filepath"
	"time"

	"github.com/lucasjo/porgex-go/config"
	"gopkg.in/mgo.v2"
)

type Connector struct {
	mSession *mgo.Session
	mInfo    *mgo.DialInfo
}

func NewConnction() Connector {

	var conn Connector
	root, _ := os.Getwd()
	configpath := filepath.Join(root, "config.gcfg")

	config, configerr := config.NewConfig(configpath)

	if configerr != nil {
		panic(configerr)
	}

	conn.mInfo = &mgo.DialInfo{
		Addrs:    []string{config.Mongo.Host + ":" + config.Mongo.Port},
		Timeout:  60 * time.Second,
		Database: config.Mongo.Databasename,
		Username: config.Mongo.Username,
		Password: config.Mongo.Password,
	}

	session, err := mgo.DialWithInfo(conn.mInfo)
	if err != nil {
		panic(err)
	}
	conn.mSession = session.Copy()

	return conn
}

func (conn *Connector) Session() *mgo.Session {

	return conn.mSession
}

func (conn *Connector) Info() *mgo.DialInfo {
	return conn.mInfo
}
