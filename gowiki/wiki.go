package main

import (
	"awesomeProject/gowiki/controller"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func defaultHandler(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/edit/FrontPage", http.StatusFound)
	fmt.Fprintf(writer, "Hi there, I lovae %s!", request.URL.Path[1:])
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		validate := validPath.FindStringSubmatch(request.URL.Path)
		if validate == nil {
			http.NotFound(writer, request)
		}

		fn(writer, request, validate[2])
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	wikiController := controller.WikiController{}
	http.HandleFunc("/view/", makeHandler(wikiController.ViewHandler))
	http.HandleFunc("/edit/", makeHandler(wikiController.EditHandler))
	http.HandleFunc("/save/", makeHandler(wikiController.SaveHandler))
	http.HandleFunc("/", defaultHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
