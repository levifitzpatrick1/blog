package templates

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var AppTemplates *template.Template

func LoadTemplates(templateDir string) {
	var err error
	AppTemplates = template.New("")

	tmplFiles, _ := filepath.Glob(filepath.Join(templateDir, "*.html"))

	AppTemplates, err = AppTemplates.ParseFiles(tmplFiles...)
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
}
func RenderTemplate(w http.ResponseWriter, tmplName string, data any) {
	buf := new(bytes.Buffer)
	err := AppTemplates.ExecuteTemplate(buf, tmplName, data)
	if err != nil {
		log.Printf("Error executing template %s: %v", tmplName, err)
		http.Error(w, fmt.Sprint("Failed to render template...", tmplName), http.StatusInternalServerError)
		return
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Printf("Error writing template %s to response: %v", tmplName, err)
	}
}
