package hellosvc

import (
	"github.com/jaijiv/go-social-api/database"
	"gopkg.in/mgo.v2/bson"
)

func Post(user string) error {
	ms := database.MongoSession()
	defer ms.Session.Close()

	col := ms.HelloCol()
	col.Insert(bson.M{"user": user})

	return nil
}

func GetAll() error {
	ms := database.MongoSession()
	defer ms.Session.Close()

	//collection := sessionCopy.DB(environment.ConnectionInfo.Database).C("hello")

	return nil
}
