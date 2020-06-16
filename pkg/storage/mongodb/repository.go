package mongodb

import (
	"go-skeleton/pkg/adding"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Storage store data in database
type Storage struct {
	s *mgo.Session
}

//AddCreator saves creator to repository
func (ss *Storage) AddCreator(c adding.Creator) error {
	var err error
	err = ss.s.DB("database-name").C("collection-name").Insert(bson.M{})
	return err
}
