package main

import (
	"fmt"
	"html/template"
	"net/http"
	//"io/ioutil"
	"./utils"
	"log"
	"os"
	"regexp"
	"strings"
)

var ERROR_501 = "<html><head><title>Error</title></head><body><h1> 500 - Internal Server Error </h1></body>"

//Entry Point
func main() {
	// set up the logger
	var logger, err = os.OpenFile("log/log.txt", os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	log.SetOutput(logger)
	http.ListenAndServe(":8080", New(logger))
}

//This function returns our object
func New(logger *os.File) http.Handler {
	var router = http.NewServeMux()
	var app = App{router, logger}
	router.HandleFunc("/", app.Index)
	router.HandleFunc("/static/", app.Static)
	router.HandleFunc("/get_summoner", app.GetSummoner)
	router.HandleFunc("/feedback", app.Feedback)
	router.HandleFunc("/get_match_info", app.GetMatchInfo)
	return app
}

//The main app object that holds the methods for the router and the actual router
type App struct {
	router *http.ServeMux
	logger *os.File
}

//Implementing the Interface method
//NOTE> This is a non-pointer receiver. The rest of the methods are pointers
func (app App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.router.ServeHTTP(w, r)
}

//Method responsible for the main page
func (app *App) Index(w http.ResponseWriter, r *http.Request) {
	var IndexTemplate, err = template.New("IndexTemplate").ParseFiles("templates/base.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else {
		IndexTemplate.Execute(w, nil)
	}
}

//the function that serves CSS
//NOTE > need to refactor and improve it to serve any kind of static files.
//NOTE > On golangs page it says that if this can take user input, then I should also make sure that the input is properly configured.
func (app *App) Static(w http.ResponseWriter, r *http.Request) {
	var path = strings.Trim(r.URL.Path, "/")
	http.ServeFile(w, r, path)
}

func (app *App) GetSummoner(w http.ResponseWriter, r *http.Request) {
	//if we do not parse the form we will get an empty map
	r.ParseForm()
	//Verify that the summoner name is a valid name : https://developer.riotgames.com/getting-started.html
	var ok, _ = regexp.Match("^[0-9\\p{L} _]+$", []byte(r.Form["SummonerName"][0]))
	if !ok {
		app.Index(w, r)
	} else {
		var summoner, err = utils.GetSummonerByName(r.Form["SummonerName"][0], r.Form["Server"][0])
		if err != nil {
			log.Println(fmt.Sprintf("FATAL : %s", err))
			app.Index(w, r)
		} else {
			var SummonerTemplate, err = template.New("SummonerTemplate").ParseFiles("templates/summoner.html")
			if err != nil {
				log.Println(fmt.Sprintf("FATAL : %s", err))
				app.Index(w, r)
			} else {
				SummonerTemplate.Execute(w, summoner)
			}
		}
	}
}

func (app *App) Feedback(w http.ResponseWriter, r *http.Request) {
	var FeedbackTemplate, err = template.New("Feedback").ParseFiles("templates/feedback.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else {
		FeedbackTemplate.Execute(w, nil)
	}
}

func (app *App) GetMatchInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var Details, Err = utils.GetMatchById(r.Form["MatchId"][0], r.Form["Server"][0])
	if Err != nil {
		w.Write([]byte("No Info"))
	} else {
		var MatchTemplate, TemplateError = template.New("MatchInfo").ParseFiles("templates/MatchInfo.html")
		if TemplateError != nil {
			fmt.Println(TemplateError)
			w.Write([]byte("No info"))
		} else {
			MatchTemplate.Execute(w, Details)
		}
	}
}
