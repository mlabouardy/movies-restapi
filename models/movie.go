package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Movie struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	CoverImage  string        `json:"cover_image" bson:"cover_image"`
	ReleaseDate time.Time     `json:"release_date" bson:"release_date"`
	Description string        `json:"description" bson:"description"`
}
