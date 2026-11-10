package main

import "slices"

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

func search_contact_by_id(id int) *Contact {
	for i := range contactList {
		if contactList[i].ID == id {
			return &contactList[i]
		}
	}
	return nil
}

func delete_contact_by_id(id int) bool {
	for i := range contactList {
		if contactList[i].ID == id {
			contactList = slices.Delete(contactList, i, i+1)
			return true
		}
	}
	return false
}