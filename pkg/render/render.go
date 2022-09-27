package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tc = make(map[string]*template.Template)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error during parsing", err)
		return
	}
}

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	_, inMap := tc[t]
	if !inMap {
		log.Println("Creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("using cached template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, tmpl)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.html",
	}

	// Parse the files
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// Add the template to the cache
	tc[t] = tmpl

	return nil
}
