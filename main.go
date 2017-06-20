package main

import ( 
		"html/template"
		"fmt"
		"net/http"
		//"io/ioutil"
		//"os"
		"strings"
		)

//Entry Point
func main() {
	http.ListenAndServe(":8080",New())
}

//This function returns our object
func New() http.Handler {
	var router = http.NewServeMux()
	var app = App{router}
	router.HandleFunc("/", app.Index)
	router.HandleFunc("/static/css/", app.Static)
	router.HandleFunc("/get_summoner", app.GetSummoner)
	return app
}

//The main app object that holds the methods for the router and the actual router
type App struct { 
	router *http.ServeMux
}

//Implementing the Interface method
//NOTE> This is a non-pointer receiver. The rest of the methods are pointers
func (app App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.router.ServeHTTP(w,r)
}

//Method responsible for the main page
func (app *App) Index(w http.ResponseWriter, r *http.Request){
	var IndexTemplate,err = template.New("IndexTemplate").ParseFiles("templates/base.html")
	fmt.Println(r.URL.Path)
	if err != nil {
		w.Write([]byte("Nope"))
	} else {
		IndexTemplate.Execute(w,nil)
	}
}

//the function that serves CSS
//NOTE > need to refactor and improve it to serve any kind of static files. 
//NOTE > On golangs page it says that if this can take user input, then I should also make sure that the input is properly configured.
func (app *App) Static(w http.ResponseWriter, r *http.Request) {
	var path = strings.Trim(r.URL.Path, "/")
	http.ServeFile(w,r,path)
}

func (app *App) GetSummoner(w http.ResponseWriter, r *http.Request) {
	//if we do not parse the form we will get an empty map
	r.ParseForm()
}