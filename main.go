package main

import (
	"net/http"
	"strconv"
	"strings"
	//"strconv"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("GET /contacts", get_contacts_handler)
	http.HandleFunc("GET /contacts/new", get_add_contact_handler)
	http.HandleFunc("GET /contacts/{id}/view", get_contacts_datails_handler)
	http.HandleFunc("GET /contacts/{id}/edit", get_edit_contact_handler)
	http.HandleFunc("GET /contacts/{id}/email", get_validate_email_handler)

	http.HandleFunc("POST /contacts/new", post_add_contact_handler)
	http.HandleFunc("POST /contacts/{id}/edit", post_edit_contact_handler)

	http.HandleFunc("DELETE /contacts/{id}", delete_contact_handler)

	http.ListenAndServe(":8080", nil)
}

func get_contacts_handler(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("q")
	page_string := r.URL.Query().Get("page")
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

	page := 1
	page_size := 10
	if page_string != "" {
		page_number, err := strconv.Atoi(page_string)
		if err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}
		page = page_number
	}

	start := (page-1)*page_size
	end := start + page_size

	if start > len(list) {
		start = len(list)
	}

	if end > len(list) {
		end = len(list)
	}

	page_contacts := list[start:end]

	plantilla := layout(contacts(page_contacts, search, page))
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

func get_validate_email_handler(w http.ResponseWriter, r *http.Request) {
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

	contact_to_validate := *contact
	contact_to_validate.Email = r.URL.Query().Get("email")

	email_error := validate_email(id_number, contact_to_validate.Email)

	w.Write([]byte(email_error))
}