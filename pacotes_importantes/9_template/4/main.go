package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Cursos{
			Curso{"Go", 40}, 
			Curso{"Python", 40}, 
			Curso{"Java", 120},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
