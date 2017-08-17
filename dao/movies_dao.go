package dao

import (
	"log"

	. "github.com/mlabouardy/movies-restapi/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MoviesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "movies"
)

func (m *MoviesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *MoviesDAO) FindAll() []Movie {
	var movies []Movie
	if err := db.C(COLLECTION).Find(bson.M{}).All(&movies); err != nil {
		log.Fatal(err)
	}
	return movies
}

func (m *MoviesDAO) FindById(id string) Movie {
	var movie Movie
	if err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie); err != nil {
		log.Fatal(err)
	}
	return movie
}

func (m *MoviesDAO) Insert(movie Movie) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

func (m *MoviesDAO) Delete(movie Movie) {
	if err := db.C(COLLECTION).Remove(&movie); err != nil {
		log.Fatal(err)
	}
}

func (m *MoviesDAO) Update(movie Movie) {
	if err := db.C(COLLECTION).UpdateId(movie.ID, &movie); err != nil {
		log.Fatal(err)
	}
}
