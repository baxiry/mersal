package main

import (
	"fmt"
	"mersal/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

// signup sing up new user handler
func Signup(c echo.Context) error {

	email := c.FormValue("email")
	pass := c.FormValue("password")

	err := insertUser(email, pass)
	if err != nil {
		fmt.Println(err)
		return c.Render(200, "sign.html", err.Error())
	}
	return c.Redirect(http.StatusSeeOther, "/login") // 303 code
}

// insertUser register new user in db
func insertUser(email, pass string) error {

	sts := "INSERT INTO users(email, password) VALUES ( ?, ?)"
	_, err := db.Exec(sts, email, pass)

	// if there is an error inserting, handle it
	if err != nil {
		return err
	}
	// be careful deferring Queries if you are using transactions
	return nil
}

func Login(c echo.Context) error {
	femail := c.FormValue("email")
	fpass := c.FormValue("password")
	userid, username, pass := selectUser(femail)
	fmt.Println("login with ", userid, username, pass)

	if pass == fpass && pass != "" {

		auth.NewSession(c, username, userid)
		//auth.NewSession(c, userid)
		return c.Redirect(http.StatusSeeOther, "/") // 303 code
		// TODO redirect to latest page
	}
	// TODO flush this message

	data := make(map[string]interface{}, 2)
	data["username"] = username
	data["userid"] = userid
	data["message"] = "username or pass is wrong!"

	fmt.Println(c.Render(200, "login.html", data))
	return nil
}

// select User info
func selectUser(femail string) (int, string, string) {
	var username, password string
	var userid int
	err := db.QueryRow(
		"SELECT id, username, password FROM users WHERE email = ?",
		femail).Scan(&userid, &username, &password)
	if err != nil {
		fmt.Println("select user ERROR: ", err.Error())
		return -1, "", ""
	}
	return userid, username, password
}

func SignPage(c echo.Context) error {
	username, userid, err := auth.GetSession(c)
	if err != nil {
		fmt.Println("SignPage error is : ", err)
	}
	data := make(map[string]interface{}, 2)
	data["username"] = username
	data["userid"] = userid

	return c.Render(200, "sign.html", data)
}

func LoginPage(c echo.Context) error {
	username, userid, err := auth.GetSession(c)
	if err != nil {
		fmt.Println("LoginPage error is : ", err)
	}
	data := make(map[string]interface{}, 2)
	data["username"] = username
	data["userid"] = userid

	fmt.Println(c.Render(200, "login.html", data))
	return nil
}
