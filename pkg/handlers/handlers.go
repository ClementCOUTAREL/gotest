package handlers

import (
	"net/http"

	"github.com/ClementCOUTAREL/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.html")
}
