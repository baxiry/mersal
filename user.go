package main

import (
	"fmt"
	"mersal/auth"
	"mersal/helps"
	"net/http"
	"strconv"
	"strings"

	"github.com/blockloop/scan"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// PutProfile updates user info contorller
func PutProfile(c echo.Context) error {
	fmt.Println("PutProfile")

	data := c.FormValue("data")
	colomn := c.FormValue("colomn")
	id := c.FormValue("userid")
	fmt.Println("user id", id)

	sess, err := session.Get("session", c)
	fmt.Println("session userid: ", sess.Values["userid"])

	userid, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("A new error:", err)
	}

	//fmt.Println("data: ", data)
	//fmt.Println("colomn: ", colomn)

	err = updateUserInfo(data, colomn, userid)
	if err != nil {
		fmt.Println("\n\n\nerror is:", err)
		return err // c.Render(200, "sign.html", "wrrone")
	}
	// return c.Redirect(http.StatusSeeOther, "/login")
	return c.String(http.StatusOK, "update profile success")
}

// acount render profile of user. ok
func acount(c echo.Context) error {
	fmt.Println("at acount func")

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
	fmt.Println("at updateAcountInfo func")

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

	// TODO update session information
	// TODO	setSession(c, username, uid.(int))

	// redirect to acoun page
	userid := strconv.Itoa(uid.(int))

	return c.Redirect(303, "/acount/"+userid)
}

// updateAcount updates Acount information
func updateAcount(c echo.Context) error {
	fmt.Println("updateAcount func")

	data := make(map[string]interface{}, 1)

	fmt.Println("start get update acount page")

	userid, _, err := auth.GetSession(c)
	helps.PrintError("error when GetSession() ", err)

	if userid == "" {
		// login first
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}

	id, _ := strconv.Atoi(userid)
	data["username"], data["email"], data["linkavatar"] = GetUserInfo(id)

	data["userid"] = userid

	fmt.Println(data)
	fmt.Println("userid is ", data["userid"])

	return c.Render(200, "upacount.html", data)
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.Render(http.StatusOK, "user.html", id)
}

func updateUserInfo(name, email string, uid int) error {

	//Update db
	stmt, err := db.Prepare("update users set username=?, email=? where id=?")
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
func GetUserInfo(userid int) (string, string, string) {
	var name, email, avatar string
	err := db.QueryRow(
		"SELECT username, email, linkavatar FROM users WHERE id = ?",
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
		"SELECT id, username, email, password FROM mersal.users WHERE email = ?",
		femail).Scan(&userid, &name, &email, &password)
	if err != nil {
		fmt.Println("no result or", err.Error())
	}
	return userid, name, email, password
}

// getUserIfor select * by Id
func getUserInfo(userid int) (user User) {

	rows, err := db.Query("SELECT * FROM users WHERE id = ?", userid)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	err = scan.Row(&user, rows)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	photo := strings.Split(user.Photos, "; ")
	user.Photos = photo[0]

	return user
}

func Profile(c echo.Context) error {
	username, userid, err := auth.GetSession(c)
	if err != nil {
		//println(c.Redirect(http.StatusSeeOther, "/login"))
		fmt.Println("no session", err)
	}

	fmt.Println("username is : ", username)
	fmt.Println("userid is : ", userid)

	id, _ := strconv.Atoi(c.Param("id"))

	data := make(map[string]interface{}, 1)
	data["user"] = getUserInfo(id)

	if id == userid {
		data["owner"] = "ok"
	}
	// for session
	data["userid"] = userid
	data["username"] = username

	fmt.Println(c.Render(200, "user.html", data))
	return nil
}

// updatePage update Page info
func UpdatePage(c echo.Context) error {

	username, userid, err := auth.GetSession(c)
	if err != nil {
		println(c.Redirect(http.StatusSeeOther, "/login"))
		fmt.Println("error of upacount handler is ", err)
		return nil
	}

	data := make(map[string]interface{}, 1)

	data["username"] = username
	data["userid"] = userid
	data["user"] = getUserInfo(userid)

	user := data["user"].(User)
	fmt.Println(user.Id)

	fmt.Println(data)

	fmt.Println(c.Render(200, "upacount.html", data))
	return nil
}

// update updates Acount information
func Update(c echo.Context) error {
	fmt.Println("update account")

	_, userid, err := auth.GetSession(c)
	if err != nil {
		println(c.Redirect(http.StatusSeeOther, "/login"))
		return nil
	}

	stmt, err := db.Prepare(
		`update users set username=? where id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute
	stmt.Exec(c.FormValue("username"), userid)

	fmt.Println(c.Redirect(http.StatusSeeOther, "/user/"+strconv.Itoa(userid)))
	return nil
}

// update user info in db
func UpdateUserInfo(field string, userid int) error {

	//Update db
	stmt, err := db.Prepare("update users set " + field + "=? where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// execute
	_, err = stmt.Exec(field, userid) // ignore affected _
	if err != nil {
		return err
	}
	return nil
}
