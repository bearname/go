package controller

import (
	m "awesomeProject/gowiki/model"
	"fmt"
	_ "fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
)

type WikiController struct {
}

func (c WikiController) ViewHandler(writer http.ResponseWriter, request *http.Request, title string) {

	p, err := c.loadPage(title)
	if err != nil {
		http.Redirect(writer, request, "/edit/"+title, http.StatusFound)
		return
	}
	c.renderTemplate(writer, "view", p)
}

func (c WikiController) EditHandler(writer http.ResponseWriter, request *http.Request, title string) {
	p, err := c.loadPage(title)
	if err != nil {
		p = &m.Page{Title: title}
	}

	c.renderTemplate(writer, "edit", p)
}

func (c WikiController) SaveHandler(writer http.ResponseWriter, request *http.Request, title string) {
	body := request.FormValue("body")
	p := &m.Page{Title: title, Body: []byte(body)}
	err := p.Save()

	c.tryInternalServerError(writer, err)

	http.Redirect(writer, request, "/view/"+title, http.StatusFound)
}

func (c WikiController) renderTemplate(writer http.ResponseWriter, tmpl string, page *m.Page) {
	files, err := template.ParseFiles("template/" + tmpl + ".html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = files.Execute(writer, page)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c WikiController) loadPage(title string) (*m.Page, error) {
	filename := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)

	search := regexp.MustCompile("\\[([a-zA-Z]+)]")

	body = search.ReplaceAllFunc(body, func(s []byte) []byte {
		group := search.ReplaceAllString(string(s), `$1`)

		fmt.Println(group)

		newGroup := "<a href='/view/" + group + "'>" + group + "</a>"
		return []byte(newGroup)
	})
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))
	return &m.Page{Title: title, Body: body}, nil
}

func (c WikiController) tryInternalServerError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
