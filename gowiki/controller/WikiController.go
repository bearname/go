package controller

import (
	m "awesomeProject/gowiki/model"
	_ "fmt"
	"html/template"
	"io/ioutil"
	"net/http"
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

func (c WikiController) loadPage(title string) (*m.Page, error) {
	filename := "data/" + title + ".txt"
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
