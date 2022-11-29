package Hangmanweb

import (
	"html/template"
	"net/http"
)

type DataTest struct {
	Name   string
	Number int
	Form   string
	Class  string
}

func BetaServe() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("index.html")
		data := DataTest{Name: "hangman", Number: 9}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/froggy/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("froggy/froggy.html")
		data := DataTest{Name: "hangman", Number: 9}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/froggy/game/win", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("empereur/win.html")
		data := DataTest{Name: "hangman", Number: 9}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/froggy/game/lost", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("empereur/lost.html")
		data := DataTest{Name: "hangman", Number: 9}
		tmpl.Execute(w, data)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
