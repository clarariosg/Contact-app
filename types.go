package main

type Contact struct {
	ID int
 	First string
	Last string
	Email string
	Phone string
} 

var contactList = []Contact {
	{First: "Ana", Last: "", Email: "ana@example.com", Phone: ""},
	{First: "Luis", Last: "", Email: "luis@example.com", Phone: ""},
}