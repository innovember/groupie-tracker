package internal

import (
	"html/template"
	"net/http"
)

var (
	templates      *template.Template
	AsciiArtResult string
)

func init() {
	templates = template.Must(template.ParseGlob("./templates/*.html"))
}
