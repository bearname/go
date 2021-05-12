package main

import (
	"awesomeProject/gowiki/controller"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I lovae %s!", r.URL.Path[1:])
}

func main() {
	{
		wikiController := controller.WikiController{}
		http.HandleFunc("/view/", wikiController.ViewHandler)
		http.HandleFunc("/edit/", wikiController.EditHandler)
		http.HandleFunc("/save/", wikiController.SaveHandler)
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
	//{
	//	page1 := &m.Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	//	err := page1.Save()
	//	if err != nil {
	//		fmt.Println("Error" + err.Error())
	//	}
	//	p2, _ := loadPage("TestPage")
	//	fmt.Println(string(p2.Body))
	//}
}
