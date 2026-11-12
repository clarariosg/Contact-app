package main

import (
	"contact-app/archiver"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("GET /contacts", get_contacts_handler)
	http.HandleFunc("GET /contacts/new", get_add_contact_handler)
	http.HandleFunc("GET /contacts/{id}/view", get_contacts_details_handler)
	http.HandleFunc("GET /contacts/{id}/edit", get_edit_contact_handler)
	http.HandleFunc("GET /contacts/{id}/email", get_validate_email_handler)
	http.HandleFunc("GET /contacts/count", get_contacts_count_handler)
	http.HandleFunc("GET /contacts/archive", get_archive_status_handler)
	http.HandleFunc("GET /contacts/archive/file", get_archive_content_handler)

	http.HandleFunc("POST /contacts/new", post_add_contact_handler)
	http.HandleFunc("POST /contacts/{id}/edit", post_edit_contact_handler)
	http.HandleFunc("POST /contacts/archive", post_start_archive_handler)

	http.HandleFunc("DELETE /contacts/{id}", delete_contact_handler)
	http.HandleFunc("DELETE /contacts", delete_all_contact_handler)
	http.HandleFunc("DELETE /contacts/archive", delete_contact_archive_handler)

	http.ListenAndServe(":8080", nil)
}

func get_contacts_handler(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("q")
	page_string := r.URL.Query().Get("page")
	
	var list []Contact

	if search != "" {
		// Itera por la lista de contactos y encuentra los que contienen el término de búsqueda en primer
		// nombre, apellido, email o telefono
		for _, c := range contactList {
			if strings.Contains(strings.ToLower(c.First), strings.ToLower(search)) ||
			strings.Contains(strings.ToLower(c.Last), strings.ToLower(search)) ||
			strings.Contains(strings.ToLower(c.Email), strings.ToLower(search)) ||
			strings.Contains(strings.ToLower(c.Phone), strings.ToLower(search)) {
				list = append(list, c)
			}
		}
	} else {
		list = contactList
	}

	page := 1
	// Si el query param de page no está vacio, se castea a int y si no da error se usa para paginacion
	if page_string != "" {
		page_number, err := strconv.Atoi(page_string)
		if err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}
		page = page_number
	}

	page_contacts := paginate_contacts(list, page, 10)

	if r.Header.Get("HX-Trigger") == "search" {
		plantilla := rows(page_contacts)
		err := plantilla.Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}
		return
	}

	plantilla := layout(contacts(page_contacts, search, page, archiver.Get()))
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

func get_contacts_details_handler(w http.ResponseWriter, r *http.Request) {
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
	if r.Header.Get("HX-Trigger") == "delete-btn" {
		w.Header().Set("HX-Redirect", "/contacts")
		w.WriteHeader(http.StatusSeeOther)
		return
	}
	
	w.WriteHeader(http.StatusOK)
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

func get_contacts_count_handler(w http.ResponseWriter, r *http.Request) {
	count := len(contactList)
	fmt.Fprintf(w, "(%d total Contacts)", count)
}

func delete_all_contact_handler(w http.ResponseWriter, r *http.Request) {
	var ids_a_eliminar []int
	ids_query := r.URL.Query()["selected_contact_ids"]
	
	for _, id := range ids_query {
		id_to_number, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid contact id", http.StatusBadRequest)
			return
		}
		ids_a_eliminar = append(ids_a_eliminar, id_to_number)
	}

	for _, contact_id := range ids_a_eliminar {
		deleted := delete_contact_by_id(contact_id)
		if !deleted {
			http.NotFound(w, r)
			return
		}
	}
	page := 1
	
	page_contacts := paginate_contacts(contactList, page, 10)

	plantilla := layout(contacts(page_contacts, "", 1, archiver.Get()))
	err := plantilla.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}

func post_start_archive_handler(w http.ResponseWriter, r *http.Request) {
	archiver := archiver.Get()
	archiver.Run()

	plantilla := archive_ui(archiver)
	err := plantilla.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}

func get_archive_status_handler(w http.ResponseWriter, r *http.Request) {
	archiver := archiver.Get()
	plantilla := archive_ui(archiver)
	err := plantilla.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}

func get_archive_content_handler(w http.ResponseWriter, r *http.Request) {
	archiver := archiver.Get()
	data := archiver.Archive_file(contactList)
	if data == nil {
		http.Error(w, "Failed to archive contact list", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Disposition", "attachment; filename=\"contacts.json\"")
	_, err := w.Write(data.([]byte))
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}

func delete_contact_archive_handler(w http.ResponseWriter, r *http.Request) {
	archiver := archiver.Get()
	archiver.Reset()
	plantilla := archive_ui(archiver)
	err := plantilla.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}