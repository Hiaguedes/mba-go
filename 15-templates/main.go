package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
	Descricao 	 string
}

type Cursos []Curso

func main() {
	var cursos Cursos = Cursos{
		{ Nome: "GO", CargaHoraria: 360, Descricao: "Seja expert em go" },
		{ Nome: "Python", CargaHoraria: 12, Descricao: "Seja expert em Python"},
		{ Nome: "Typescript", CargaHoraria: 6, Descricao: "Seja expert em Typescript" },
		{ Nome: "Cobol", CargaHoraria: 10000, Descricao: "Seja expert em Cobol" },
	}
	// tmp := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horaria: {{.CargaHoraria}}"))

	tmp := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := tmp.Execute(os.Stdout, cursos)

	if err != nil {
		panic(err)
	}
}