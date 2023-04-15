package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Cursos{
		{"Curso de Go", 40},
		{"Curso de Python", 20},
		{"Curso de Java", 30},
	})
	if err != nil {
		panic(err)
	}

}
