package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "github.com/mlabouardy/movies-restapi/config"
	. "github.com/mlabouardy/movies-restapi/dao"
	. "github.com/mlabouardy/movies-restapi/models"
)

var config = Config{}
var dao = MoviesDAO{}

func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	var movies = dao.FindAll()
	json.NewEncoder(w).Encode(&movies)
}

func FindMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var movie = dao.FindById(params["id"])
	json.NewEncoder(w).Encode(&movie)
}

func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}
	movie.ID = bson.NewObjectId()
	if err := dao.Insert(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJson(w, http.StatusCreated, movie)
}

func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies", CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies/:id", FindMovieEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
