package main

import (
	"net/http"
	//"github.com/a-h/templ"
)

func main() {
	http.HandleFunc("GET /contacts", get_contacts_handler)

	http.ListenAndServe(":8080", nil)
}

func get_contacts_handler(w http.ResponseWriter, r *http.Request) {
	plantilla := contacts()
	plantilla.Render(r.Context(), w)
}
