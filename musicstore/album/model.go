package album

import (
	"gopkg.in/mgo.v2/bson"
)

// Album ...
type Album struct {
	ID     bson.ObjectId `bson:"_id"`
	Title  string        `json:"title"`
	Artist string        `json:"artist"`
	Year   string        `json:"year"`
}

// Albums ...
type Albums []Album
