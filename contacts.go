package main

import "fmt"

// get an username by email
func getAllContacts(femail string) (int, string, string, string) {
	var name, email, password string
	var userid int
	err := db.QueryRow(
		"SELECT userid, username, email, password FROM comments.users WHERE email = ?",
		femail).Scan(&userid, &name, &email, &password)
	if err != nil {
		fmt.Println("no result or", err.Error())
	}
	return userid, name, email, password
}

type Contact struct {
	contactid   int
	contactname string
}

// get an username by email
func getMyContcts(userid int) (contacts []Contact) {
	contact := Contact{}
	err := db.QueryRow(
		"SELECT userid, username FROM mersal.users", //" WHERE email = ?",
	).Scan(&contact.contactid, &contact.contactname)
	if err != nil {
		fmt.Println("no result or", err.Error())
	}
	return contacts
}
