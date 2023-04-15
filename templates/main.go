package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {

	curso := Curso{"Curso de Go", 40}
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("O curso {{.Nome}} tem carga horária de {{.CargaHoraria}} horas.")
	err := tmp.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

}
