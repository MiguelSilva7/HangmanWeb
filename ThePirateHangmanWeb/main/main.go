package main

import (
	"fmt"
	"hangman"
	"log"
	"net/http"
	"text/template"
)

func Page_Home(rw http.ResponseWriter, r *http.Request, Pts *hangman.HangData) {
	template, err := template.ParseFiles("./index.html", "./template/header.html", "./template/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(rw, Pts)
}

func Page_Hangman(rw http.ResponseWriter, r *http.Request, Pts *hangman.HangData) {
	template, err := template.ParseFiles("./pages/hangman.html", "./template/header.html", "./template/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(rw, Pts)
}

func Page_Infos(w http.ResponseWriter, r *http.Request, Pts *hangman.HangData) {
	template, err := template.ParseFiles("./pages/infos.html", "./template/header.html", "./template/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, Pts)
}

func InitialiseStruct(Pts *hangman.HangData) {
	Pts.Life = 10
	Pts.To_found = hangman.WordSelector()
	Pts.Founded = hangman.Founded(Pts.To_found)
	Pts.To_found_RuneVersion = hangman.StringToSliceRune(Pts.To_found)
	Pts.Correct = false
	Pts.Founded_RuneVersion = hangman.SliceRuneToString(Pts.Founded)
}

func main() {
	//====== WEB SERVER
	fmt.Printf("Starting server at port 8080\n")

	Pts := &hangman.HangData{}
	InitialiseStruct(Pts)

	fs := http.FileServer(http.Dir("styles/"))
	http.Handle("/styles/", http.StripPrefix("/styles/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Page_Home(w, r, Pts)
	})

	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		Page_Hangman(w, r, Pts)
	})

	http.HandleFunc("/infos", func(w http.ResponseWriter, r *http.Request) {
		Page_Infos(w, r, Pts)
	})

	http.ListenAndServe(":8080", nil)
	hangman.Play(Pts)
	//======

	//fmt.Println(tab_words[nbr_words], "<--//===============TEST=imprime le mot du jeu
	//fmt.Println(to_found_RuneVersion, "<--//===============TEST=mot à trouver version rune

	// Fonction principal du jeu (loop)
	InitialiseStruct(Pts)
	hangman.Play(Pts)
}
