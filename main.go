package main

import (
	"net/http"
	"strconv"
	"strings"
	//"strconv"
)

func main() {
	http.HandleFunc("GET /contacts", get_contacts_handler)
	http.HandleFunc("GET /contacts/new", get_add_contact_handler)
	http.HandleFunc("GET /contacts/{id}/view", get_contacts_datails_handler)
	http.HandleFunc("GET /contacts/{id}/edit", get_edit_contact_handler)

	http.HandleFunc("POST /contacts/new", post_add_contact_handler)
	http.HandleFunc("POST /contacts/{id}/edit", post_edit_contact_handler)
	http.HandleFunc("POST /contacts/{id}/delete", delete_contact_handler)

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

func get_add_contact_handler(w http.ResponseWriter, r *http.Request) {
	var contact Contact
	plantilla := layout(new(contact))
	err := plantilla.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}

func post_add_contact_handler(w http.ResponseWriter, r *http.Request) {
	first := r.FormValue("first_name")
	last := r.FormValue("last_name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")

	var contact Contact

	contact.ID = len(contactList) + 1
	contact.First = first
	contact.Last = last
	contact.Email = email
	contact.Phone = phone

	contactList = append(contactList, contact)
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func get_contacts_datails_handler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	id_number, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	contact := search_contact_by_id(id_number)
	if contact == nil {
		http.NotFound(w, r)
		return
	}

	plantilla := layout(view(*contact))
	err1 := plantilla.Render(r.Context(), w)
	if err1 != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}

func get_edit_contact_handler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	id_number, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	contact := search_contact_by_id(id_number)
	if contact == nil {
		http.NotFound(w, r)
		return
	}
	plantilla := layout(edit(*contact))
	err1 := plantilla.Render(r.Context(), w)
	if err1 != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}

func post_edit_contact_handler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	id_number, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	contact := search_contact_by_id(id_number)
	if contact == nil {
		http.NotFound(w, r)
		return
	}

	contact.First = r.FormValue("first_name")
	contact.Last = r.FormValue("last_name")
	contact.Phone = r.FormValue("phone")
	contact.Email = r.FormValue("email")

	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func delete_contact_handler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	id_number, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	deleted := delete_contact_by_id(id_number)
	if deleted == false {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}