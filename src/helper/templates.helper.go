package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var listeTemplate *template.Template

func Load() {

	temp, tempErr := template.ParseGlob("../templates/*.html")
	if tempErr != nil {
		log.Fatalf("Erreur chargement templates - %s", tempErr.Error())
		return
	}
	listeTemplate = temp
	fmt.Println("Templates chargés avec succès")
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	var buffer bytes.Buffer
	errRender := listeTemplate.ExecuteTemplate(&buffer, name, data)
	if errRender != nil {
		fmt.Println(errRender)
		http.Error(w, "Erreur de rendu template", http.StatusInternalServerError)
		return
	}
	buffer.WriteTo(w)
}