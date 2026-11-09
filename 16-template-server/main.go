package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
	Descricao 	 string
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func TemplateHandler(w http.ResponseWriter, r *http.Request) {
	var cursos Cursos = Cursos{
		{ Nome: "GO", CargaHoraria: 360, Descricao: "Seja expert em go" },
		{ Nome: "python", CargaHoraria: 12, Descricao: "Seja expert em Python"},
		{ Nome: "Typescript", CargaHoraria: 6, Descricao: "Seja expert em Typescript" },
		{ Nome: "Cobol", CargaHoraria: 10000, Descricao: "Seja expert em Cobol" },
	}
	// tmp := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horaria: {{.CargaHoraria}}"))

	tmp := template.New("template.html")
	tmp.Funcs(template.FuncMap{"ToUpper": ToUpper})
	tmp = template.Must(tmp.ParseFiles("template.html"))

	err := tmp.Execute(w, cursos)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", TemplateHandler)

	http.ListenAndServe(":8080", mux)
}