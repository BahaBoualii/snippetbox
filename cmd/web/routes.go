package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileSever := http.FileServer(neutredFileSystem{http.Dir("./ui/static")})

	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileSever))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return app.logRequest(secureHeaders(mux))
}
