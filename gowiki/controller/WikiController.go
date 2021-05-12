package main

import (
	m "awesomeProject/gowiki/model"
	_ "fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type WikiController struct {
}

func (c WikiController) ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := c.loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	c.renderTemplate(w, "view", p)
}

func (c WikiController) loadPage(title string) (*m.Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &m.Page{Title: title, Body: body}, nil
}

func (c WikiController) renderTemplate(w http.ResponseWriter, tmpl string, p *m.Page) {
	t, err := template.ParseFiles("template/" + tmpl + ".html")
	c.tryInternalServerError(w, err)
	err = t.Execute(w, p)

	c.tryInternalServerError(w, err)
}

func (c WikiController) tryInternalServerError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c WikiController) EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := c.loadPage(title)
	if err != nil {
		p = &m.Page{Title: title}
	}

	c.renderTemplate(w, "edit", p)
}

func (c WikiController) SaveHandler(writer http.ResponseWriter, request *http.Request) {
	title := request.URL.Path[len("/save/"):]
	body := request.FormValue("body")
	p := &m.Page{Title: title, Body: []byte(body)}
	err := p.Save()

	c.tryInternalServerError(writer, err)

	http.Redirect(writer, request, "/view/"+title, http.StatusFound)
}
