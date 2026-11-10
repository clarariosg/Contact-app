package main

type Contact struct {
	ID int
 	First string
	Last string
	Email string
	Phone string
} 

var contactList = []Contact {
	{ID: 1, First: "Ana", Last: "", Email: "ana@example.com", Phone: ""},
	{ID: 2, First: "Luis", Last: "", Email: "luis@example.com", Phone: ""},
}

func search_contact_by_id(id int) Contact {
	var resultado Contact
	for _, c := range contactList {
		if c.ID == id {
			return c
		}
	}
	return resultado
}