package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.errorLog.Print(err.Error())
		app.serverError(w, err)
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		app.serverError(w, err)
	}

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d ...", id)
	// w.Write([]byte("Display a specific snippet with ID %d..."))
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
