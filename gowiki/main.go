package main

import (
	"awesomeProject/gowiki/controller"
	"awesomeProject/gowiki/controller/api"
	"awesomeProject/gowiki/persistense"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view|api)/([a-zA-Z0-9]+)$")

func defaultHandler(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/edit/FrontPage", http.StatusFound)
	_, _ = fmt.Fprintf(writer, "Hi there, I lovae %s!", request.URL.Path[1:])
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		validate := validPath.FindStringSubmatch(request.URL.Path)
		if validate == nil {
			http.NotFound(writer, request)
			return
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
	bookRepository := persistense.NewSqliteBookRepository("sqlite-database.db")
	var bookController = api.NewBookController(bookRepository)

	books := bookRepository.GetBooks()
	for _, book := range books {
		fmt.Println(book.Title + " " + book.Author)
	}
	router := mux.NewRouter()

	router.HandleFunc("/view/{page}", makeHandler(wikiController.ViewHandler)).Methods(http.MethodGet)
	router.HandleFunc("/edit/{page}", makeHandler(wikiController.EditHandler)).Methods(http.MethodGet)
	router.HandleFunc("/save/", makeHandler(wikiController.SaveHandler)).Methods(http.MethodPost)
	router.HandleFunc("/", defaultHandler).Methods(http.MethodGet)

	subRouter := router.PathPrefix("/api/v1").Subrouter()
	subRouter.HandleFunc("/book/{id}", bookController.GetBook).Methods(http.MethodGet)
	subRouter.HandleFunc("/book/", bookController.ShowBooks).Methods(http.MethodGet)
	subRouter.HandleFunc("/book/create", bookController.CreateBook).Methods(http.MethodPost)

	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	log.Fatal(http.ListenAndServe(":8080", router))
}
