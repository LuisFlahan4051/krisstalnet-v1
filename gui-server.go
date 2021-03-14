package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	port   = "4051"
	urlGui = "http://localhost:" + port + "/"
)

func index(writer http.ResponseWriter, request *http.Request) {
	indexTemplate := template.Must(template.ParseFiles("build/index.html"))
	indexTemplate.Execute(writer, nil)
}

func addGUIHandler(mux *mux.Router) *mux.Router {
	staticFiles := http.FileServer(http.Dir("build/static/"))

	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", staticFiles))
	mux.HandleFunc("/", index)

	fmt.Println("\nPort " + port + " added succesfully to mux.\n")
	return mux
}

func newMux() *mux.Router {
	mux := mux.NewRouter()

	mux.Use(cors.AllowAll().Handler)

	mux = addGUIHandler(mux)
	return mux
}

func runServer(mux *mux.Router) {
	fmt.Println("Server working fine!")
	fmt.Println("Now, you can enter to " + urlGui + " in your web browser.\n")
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func main() {
	//FOR BUILD > go build -ldflags "-H windowsgui" -o main.exe

	prepareMux := newMux()

	runServer(prepareMux)
}
