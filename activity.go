package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func ActivityPage(c echo.Context) error {
	data := make(map[string]interface{}, 1)
	sess, _ := session.Get("session", c)

	uid := sess.Values["userid"]
	username := sess.Values["username"]

	data["username"] = username

	if uid == nil {
		// login first
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}

	data["user"] = getUserInfo(uid.(int))

	data["userid"] = uid

	fmt.Println(data)

	return c.Render(200, "activity.html", data)
}

// gets all user information for update this info
func MyLike(userid int) (string, string, string, string) {
	var name, email, phon, avatar string
	err := db.QueryRow(
		"SELECT username, email,phon, linkavatar FROM stores.users WHERE userid = ?",
		userid).Scan(&name, &email, &phon, &avatar)
	if err != nil {
		fmt.Println("no result or", err.Error())
	}
	return name, email, phon, avatar
}

// updateAcount updates Acount information
func LikeMe(c echo.Context) error {
	data := make(map[string]interface{}, 1)
	sess, _ := session.Get("session", c)

	uid := sess.Values["userid"]
	username := sess.Values["username"]

	data["username"] = username

	if uid == nil {
		// login first
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}

	data["user"] = getUserInfo(uid.(int))

	data["userid"] = uid

	fmt.Println(data)

	return c.Render(200, "activity.html", data)
}

// acount render profile of user.
func MyFavoriet(c echo.Context) error {
	sess, _ := session.Get("session", c)
	data := make(map[string]interface{}, 2)
	data["username"] = sess.Values["username"]
	fmt.Println("username is ", data["username"])
	data["userid"] = sess.Values["userid"]
	fmt.Println("user id or user is : ", data["userid"])
	// TODO get all info like foto from db

	if data["userid"] == nil {
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}
	return c.Render(200, "activity.html", data)
}

// acount render profile of user.
func MatchLike(c echo.Context) error {
	sess, _ := session.Get("session", c)
	data := make(map[string]interface{}, 2)
	data["username"] = sess.Values["username"]
	fmt.Println("username is ", data["username"])
	data["userid"] = sess.Values["userid"]
	fmt.Println("user id or user is : ", data["userid"])
	// TODO get all info like foto from db

	if data["userid"] == nil {
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}
	return c.Render(200, "activity.html", data)
}

// acount render profile of user.
func MatchFavorites(c echo.Context) error {
	sess, _ := session.Get("session", c)
	data := make(map[string]interface{}, 2)
	data["username"] = sess.Values["username"]
	fmt.Println("username is ", data["username"])
	data["userid"] = sess.Values["userid"]
	fmt.Println("user id or user is : ", data["userid"])
	// TODO get all info like foto from db

	if data["userid"] == nil {
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}
	return c.Render(200, "activity.html", data)
}
