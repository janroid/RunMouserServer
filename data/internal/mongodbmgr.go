package internal

import (
	"github.com/name5566/leaf/db/mongodb"
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2/bson"
)

var dialContext = new(mongodb.DialContext)

func init() {
	Connect()
}

func Connect() {
	c, err := mongodb.Dial("localhost", 10)
	if err != nil {
		log.Error("mongodbmgr.go - Connect error ! err = %v ", err)

		return
	}

	dialContext = c
	log.Debug("mongodbmgr.Connect success")
}

func Find(db string, collection string, docs interface{}, result interface{}) error {
	c := dialContext
	s := c.Ref()
	defer c.UnRef(s)

	err := s.DB(db).C(collection).Find(docs).One(result)

	if err != nil {
		log.Error("mongodbmgr.Find error ! %v", err)
	}

	return err
}

func FindAll(db string, collection string, result interface{}) error {
	c := dialContext
	s := c.Ref()
	defer c.UnRef(s)

	err := s.DB(db).C(collection).Find(bson.M{}).All(result)

	if err != nil {
		log.Error("mongodbmgr.FindAll error ! %v", err)
	}

	return err
}

func Insert(db string, collection string, docs interface{}) error {
	c := dialContext
	s := c.Ref()

	defer c.UnRef(s)

	err := s.DB(db).C(collection).Insert(docs)
	if err != nil {
		log.Error("mongodbmgr.Insert error : %v", err)
	}

	return err
}

func Update(db string, collection string, docs interface{}, result interface{}) error {
	c := dialContext
	s := c.Ref()

	defer c.UnRef(s)

	err := s.DB(db).C(collection).Update(docs, result)
	if err != nil {
		log.Error("mongodbmgr.Update error : %v", err)
	}

	return err
}
