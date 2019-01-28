package album

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct{}

func (r Repository) GetAlbums() Albums {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()

	c := session.DB(DBNAME).C(DOCNAME)
	results := Albums{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write resutls", err)
	}

	return results
}

func (r Repository) AddAlbum(album Album) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	album.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(album)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// SERVER the DB server
const SERVER = "localhost:27017"

// DBNAME the name of the DB instance
const DBNAME = "musicstore"

// DOCNAME the name of the document
const DOCNAME = "albums"
