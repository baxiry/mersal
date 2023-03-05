package main

import (
	"fmt"
	"meet/auth"
	"meet/helps"
	"net/http"
	"os"
	"strings"

	"github.com/blockloop/scan"
	"github.com/labstack/echo/v4"
)

type User struct {
	Id        int    `db:"Id"`
	Username  string `db:"Username"`
	Password  string `db:"Password"`
	Email     string `db:"Email"`
	Gender    string `db:"Gender"`
	Age       int    `db:"Age"`
	Height    int    `db:"Height"`
	Weight    int    `db:"Weight"`
	Languages string `db:"Languages"`
	Profess   string `db:"Profess"`
	Contry    string `db:"Country"`
	Descript  string `db:"Descript"`
	Photos    string `db:"Photos"`
}

// getCatigories get all photo name of catigories.
func getRecentUsers() (users []User) {
	rows, err := db.Query("SELECT Id, Username, Email, Photos, Gender from social.users;")
	defer rows.Close()
	err = scan.Rows(&users, rows)
	if err != nil {

		helps.PrintError("error from schan.Rows: ", err)
		os.Exit(1)
	}
	return users
}

// HomePage get home page
func HomePage(c echo.Context) error {

	username, userid, err := auth.GetSession(c)

	helps.PrintError("get session", err)
	println("userid is ", userid)
	println("username is ", username)

	data := make(map[string]interface{}, 3)
	data["username"] = username
	data["userid"] = userid
	users := getRecentUsers()

	//	fmt.Print("profile info", ProfileInfo(userid), "\n\n")
	for i := range users {
		photos := strings.Split(users[i].Photos, "; ")
		users[i].Photos = SetAvatar(users[i].Gender, photos[0])
		fmt.Println(users[i].Photos)
	}

	data["users"] = users
	//data["user"] = ProfileInfo(userid)
	fmt.Println(c.Render(http.StatusOK, "home.html", data))
	return nil
}

// get Profile with all info
func ProfileInfo(userid int) (profile User) {
	rows, err := db.Query("select * from users where userid = ?", userid)
	if err != nil {
		fmt.Println("\nat query func owner id db select ", err)
	}
	defer rows.Close() // ??
	err = scan.Rows(&profile, rows)
	println(err)
	return profile
}

// SetAvatar set real or blank, man or woman avatar
func SetAvatar(gen, photo string) string {
	if photo != "" {
		return photo
	}
	if gen == "m" {
		return "bman.jpg"
	}
	if gen == "f" {
		return "bwoman.jpg"
	}
	return ""
}
