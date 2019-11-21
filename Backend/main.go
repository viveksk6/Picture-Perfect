package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
	//qstring2:= qstring + " where title like %a%"
	//qstring2:= "SELECT movieId, title, summary, genre, img, language, certificate from movies where title like '%a%'"
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

func queryDbMovie(movieId string)  {
	db, err := sql.Open("mysql", "clumio:qwerty123@tcp(127.0.0.1:3306)/pictureperfect")
	if err != nil {
		panic(err.Error())
	}
	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	q := "SELECT movieId, title, summary, genre, img, language, certificate from movies where movieId = "
	qstring := q + "'" + movieId + "'"
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
	w.Header().Set("Access-Control-Allow-Origin", "*")


	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	movieDetails = nil
	pageSize = ""
	pageNo = ""
	queryValues := r.URL.Query()
	pageSize = queryValues.Get("pageSize")
	pageNo = queryValues.Get("pageNo")



	if pageNo!="" && pageSize!="" {
		pageNoInt, err:= strconv.Atoi(pageNo)
		if err != nil {
			panic(err.Error())
		}
		pageSizeInt, err:= strconv.Atoi(pageSize)
		if err != nil {
			panic(err.Error())
		}
		pageNo = strconv.Itoa(pageSizeInt*(pageNoInt-1))
		queryDb()
		fmt.Println("Endpoint Hit: returnAllMovies")
		json.NewEncoder(w).Encode(movieDetails)
	} else {
		fmt.Fprintf(w, "Invalid query parameters")
	}


}

func returnSingleMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	movieId:= vars["movieId"]
	movieDetails = nil
	queryDbMovie(movieId)
	fmt.Println("Endpoint Hit: returnSingleMovie")
	json.NewEncoder(w).Encode(movieDetails)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllMovies).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/all/{movieId}", returnSingleMovie).Methods("GET", "OPTIONS")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Go MySQL Tutorial")
	handleRequests()
}
