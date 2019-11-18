package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"

	//"github.com/davecgh/go-spew/spew"
)

type MovieDetails struct {
	MovieId   int    `json:"movieId"`
	Title string `json:"title"`
	Summary string `json:"summary"`
	Genre string `json:"genre"`
	Img string `json:"img"`
	Language string `json:"language"`
	Certificate string `json:"certificate"`
}

//var pageSizeInt int
//var movieDetails = make([]MovieDetails, pageSizeInt)
var movieDetails []MovieDetails
var pageSize, pageNo string
var err error

func queryDb()  {
	db, err := sql.Open("mysql", "clumio:qwerty123@tcp(127.0.0.1:3306)/pictureperfect")
	if err != nil {
		panic(err.Error())
	}
	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	q := "SELECT movieId, title, summary, genre, img, language, certificate from movies limit "
	qstring := q + pageNo + "," + pageSize
	results, err := db.Query(qstring)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for results.Next() {
		var md MovieDetails

		err = results.Scan(&md.MovieId, &md.Title,  &md.Summary, &md.Genre, &md.Img, &md.Language, &md.Certificate)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		movieDetails = append(movieDetails, md)
	}
}

func homePage(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
	movieDetails = nil
	queryValues := r.URL.Query()
	pageSize = queryValues.Get("pageSize")
	pageNoInt, err:= strconv.Atoi(queryValues.Get("pageNo"))
	if err != nil {
		panic(err.Error())
	}
	pageSizeInt, err:= strconv.Atoi(pageSize)
	if err != nil {
		panic(err.Error())
	}
	pageNo = strconv.Itoa(pageSizeInt*(pageNoInt-1))
	queryDb()
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(movieDetails)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Go MySQL Tutorial")
	handleRequests()
}
