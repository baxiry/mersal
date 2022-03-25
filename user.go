package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// acount render profile of user. ok
func acount(c echo.Context) error {
	sess, _ := session.Get("session", c)
	data := make(map[string]interface{}, 2)
	data["username"] = sess.Values["username"]
	data["userid"] = sess.Values["userid"]
	// TODO get all info like foto from db

	if data["userid"] == nil {
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}
	return c.Render(200, "acount.html", data)
}

// updateAcount updates Acount information
func updateAcountInfo(c echo.Context) error {
	//data := make(map[string]interface{},1)
	sess, _ := session.Get("session", c)
	uid := sess.Values["userid"]
	if uid == nil {
		// login first
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}

	username := c.FormValue("username")
	email := c.FormValue("email")
	fmt.Println("username+email is :", username, email)

	err := updateUserInfo(username, email, uid.(int))
	if err != nil {
		fmt.Println("error at update db function", err)
	}

	// update session information
	setSession(c, username, uid.(int))

	// redirect to acoun page
	userid := strconv.Itoa(uid.(int))

	return c.Redirect(303, "/acount/"+userid)
}

// updateAcount updates Acount information
func updateAcount(c echo.Context) error {
	data := make(map[string]interface{}, 1)
	sess, _ := session.Get("session", c)

	userid := sess.Values["userid"]
	username := sess.Values["username"]

	data["username"] = username

	if userid == nil {
		// login first
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}

	data["username"], data["email"], data["linkavatar"] = getUserInfo(userid.(int))

	data["userid"] = userid

	fmt.Println(data)

	return c.Render(200, "upacount.html", data)
}

//
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.Render(http.StatusOK, "user.html", id)
}

func updateUserInfo(name, email string, uid int) error {

	//Update db
	stmt, err := db.Prepare("update  comments.users set username=?, email=? where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute
	res, err := stmt.Exec(name, email, uid)
	if err != nil {
		return err
	}

	a, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println("efected foto update: ", a) // 1
	return nil
}

// gets all user information for update this info
func getUserInfo(userid int) (string, string, string) {
	var name, email, avatar string
	err := db.QueryRow(
		"SELECT username, email, linkavatar FROM comments.users WHERE userid = ?",
		userid).Scan(&name, &email, &avatar)
	if err != nil {
		fmt.Println("no result or", err.Error())
	}
	fmt.Println("name is : ", name, "email is : ", email, "avatar is ", avatar)
	return name, email, avatar
}

// get an username by email
func getUsername(femail string) (int, string, string, string) {
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
