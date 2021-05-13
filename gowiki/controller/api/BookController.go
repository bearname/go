package api

import (
	"awesomeProject/gowiki/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type BookController struct {
	repository model.BookRepository
}

func NewBookController(repository model.BookRepository) *BookController {
	controller := new(BookController)
	controller.repository = repository
	return controller
}

func (c *BookController) ShowBooks(writer http.ResponseWriter, request *http.Request) {
	books := c.repository.GetBooks()

	c.jsonResponse(writer, books)
}

func (c *BookController) CreateBook(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var t model.Book
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t.Author + " " + t.Title)
	book := c.repository.InsertBook(t.Author, t.Title)
	c.jsonResponse(writer, book)
}

func (c *BookController) jsonResponse(writer http.ResponseWriter, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = writer.Write(jsonData)
	if err != nil {
		return
	}
}

func (c *BookController) GetBook(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idString, ok := vars["id"]
	if !ok {
		fmt.Println("idString is missing in parameters")
	}
	fmt.Println(`idString := `, idString)
	var id int64
	_, err := fmt.Sscan(idString, &id)
	if err != nil {
		return
	}

	c.jsonResponse(writer, c.repository.GetBook(id))
}
