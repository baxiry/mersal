package main

import (
	"fmt"
	"meet/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

// gets all user information for update this info
func SelectMessages(userid int) (string, string, string, string) {
	var name, email, phon, avatar string
	err := db.QueryRow(
		"SELECT username, email,phon, linkavatar FROM users WHERE userid = ?",
		userid).Scan(&name, &email, &phon, &avatar)
	if err != nil {
		fmt.Println("no result or", err.Error())
	}
	return name, email, phon, avatar
}

// updateAcount updates Acount information
func updateMessage(c echo.Context) error {

	username, userid, err := auth.GetSession(c)
	if err == nil {
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}

	data := make(map[string]interface{}, 2)
	data["username"] = username
	data["userid"] = userid

	data["user"] = getUserInfo(userid)

	fmt.Println(data)

	return c.Render(200, "upacount.html", data)
}

// acount render profile of user.
func MessagesPage(c echo.Context) error {

	data := make(map[string]interface{}, 2)
	username, userid, err := auth.GetSession(c) // session.Get("session", c)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}

	data["username"] = username
	data["userid"] = userid
	fmt.Println("username is ", username)
	fmt.Println("user id or user is : ", userid)
	// TODO get all info like foto from db
	return c.Render(200, "messages.html", data)
}
