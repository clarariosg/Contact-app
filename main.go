package main

import (
	"net/http"
	"strings"
	//"github.com/a-h/templ"
	//"strconv"
)

func main() {
	http.HandleFunc("GET /contacts", get_contacts_handler)

	http.ListenAndServe(":8080", nil)
}

func get_contacts_handler(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("q")

	var list []Contact

	if search != "" {
		for _, c := range contactList {
			if strings.ToLower(c.First) == strings.ToLower(search) {
				list = append(list, c)
			}
		} 
	} else {
		list = contactList
	}
	
	plantilla := layout(contacts(list, search))
	err := plantilla.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}
