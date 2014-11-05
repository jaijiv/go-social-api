package database

import (
	"github.com/jaijiv/go-social-api/environment"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

type MgoSession struct {
	Session *mgo.Session
}

var mongoSession *mgo.Session

func ConnectMongo(connectionInfo environment.DBConnection) error {
	var err error
	/*
	 * Create the connection
	 */
	log.Println("Connecting to mongo database...")

	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{connectionInfo.Address},
		Timeout:  60 * time.Second,
		Database: connectionInfo.Database,
		Username: connectionInfo.UserName,
		Password: connectionInfo.Password,
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	mongoSession, err = mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	mongoSession.SetMode(mgo.Monotonic, true)

	return nil
}

func MongoSession() *MgoSession {
	return &MgoSession{mongoSession.Copy()}
}

func (ms *MgoSession) HelloCol() *mgo.Collection {
	return ms.Session.DB(environment.ConnectionInfo.Database).C("hello")
}
