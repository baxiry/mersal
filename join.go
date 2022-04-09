package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func insertUser(user, email, pass string) error {
	sql := `INSERT INTO mersal.users(username, email, password) VALUES (?,?,?)`

	res, err := db.Exec(sql, user, email, pass)

	if err != nil {
		fmt.Println("err is :\n ", err.Error())
		return err
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)
	return nil
}

// signup
func signup(c echo.Context) error {

	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	fmt.Println("\n\n\n", username, password, email)
	err = insertUser(username, email, password)
	if err != nil {
		fmt.Println("\n\n\nerror is:", err)
		return err // c.Render(200, "sign.html", "wrrone")
	}
	// return c.Redirect(http.StatusSeeOther, "/login")
	return c.String(http.StatusOK, "signup success")
}

// login if user info is correct
func login(c echo.Context) error {
	femail := c.FormValue("email")
	fpass := c.FormValue("password")
	userid, username, email, pass := getUsername(femail)

	if pass == fpass && femail == email {
		//userSession[email] = username
		setSession(c, username, userid)
		//return c.Redirect(http.StatusSeeOther, "/") // 303 code

		return c.String(200, "success!") // TODO redirect to latest page
	}
	//return c.Render(200, "login.html", data)
	return c.String(404, "not success!")
}

func setSession(c echo.Context, username string, userid int) {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60, // = 1h,
		HttpOnly: true,    // no websocket or any thing else
	}
	sess.Values["username"] = username
	sess.Values["userid"] = userid
	sess.Save(c.Request(), c.Response())
}

// for web
func signPage(c echo.Context) error {
	data := make(map[string]interface{}, 1)
	sess, _ := session.Get("session", c)
	data["userid"] = sess.Values["userid"]
	data["username"] = sess.Values["username"]
	return c.Render(200, "sign.html", data)
	//fmt.Println( c.Render(200, "sign.html", sess.Values["userid"].(int))); return nil
}
