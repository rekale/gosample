package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"big_projects/counts"
	"big_projects/users"

	_ "github.com/lib/pq"
	"github.com/tokopedia/sqlt"
	"github.com/urfave/negroni"
)

var view *template.Template
var db *sqlt.DB
var err error

func init() {
	dsn := "host=devel-postgre.tkpd port=5432 user=mh180102 password=CCyy8dkBXTJDPN dbname=tokopedia-user sslmode=disable"
	db, err = sqlt.Open("postgres", dsn)
	if err != nil {
		log.Println("Error ", err.Error())
	}

	if db.Ping() != nil {
		log.Println("connect database fail")
		os.Exit(1)
	} else {
		log.Println("connect database success")
	}

	view = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

	userHandler := users.NewHandler(db, view)
	mux := http.NewServeMux()

	mux.HandleFunc("/", userHandler.Home)

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(counts.IncrementsMiddleware))
	n.UseHandler(mux)

	http.ListenAndServe(":3001", n)
}
