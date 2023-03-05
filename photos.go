package main

import (
	"fmt"
	"io"
	"meet/auth"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// phootosPage router fo update Fotos Page
func PhotosPage(c echo.Context) error {

	username, userid, err := auth.GetSession(c)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login") // 303 code
	}

	data := make(map[string]interface{}, 3)
	photos := getUserFotos(userid)
	data["photos"] = photos
	data["username"] = username

	fmt.Println("photo page. user id is ", userid)
	data["userid"] = userid

	fmt.Println("fotos is : ", data)

	return c.Render(http.StatusOK, "upfotos.html", data)
}

// updateFotos updates photos
func UpPhotos(c echo.Context) error {

	pid := c.Param("id")
	fmt.Println("param id is", pid)
	id, err := strconv.Atoi(pid)
	if err != nil {
		fmt.Println("id error", err)
		return err
	}

	// from her :
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["files"]
	//fmt.Println("files is :", files[0].Filename)
	picts := ""
	for _, v := range files {
		picts += v.Filename
		picts += "; "
		// TODO Rename pictures.
	}

	// update photo link in db
	err = UpdatePhotos(picts, id)

	if err != nil {
		fmt.Println("error in update product foto", err)
	}

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			fmt.Println("err in file.Open()")
			return err
		}
		defer src.Close()

		// photoFold()
		dst, err := os.Create("../files/" + file.Filename)
		if err != nil {
			fmt.Println("err in io.Create()")
			return err
		}
		defer dst.Close()
		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			fmt.Println("err in io.Copy()")
			return err
		}
	}

	err = c.Redirect(http.StatusSeeOther, "/user/"+pid)
	if err != nil {
		fmt.Println("\nerr when update user photo", err)
	}
	return nil
}

// update fotos name in database
func UpdatePhotos(photos string, userid int) error {
	//Update db
	stmt, err := db.Prepare("update  users set photos=? where userid=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(photos, userid)
	if err != nil {
		return err
	}
	return nil
}

// selecte fotos from db
func getUserFotos(userid int) (photos []string) {
	var picts string
	err := db.QueryRow(
		"SELECT photos FROM users WHERE userid = ?",
		userid).Scan(&picts)
	if err != nil {
		return nil
	}
	list := strings.Split(picts, "; ")
	// TODO split return 2 item in some casess, is this a bug ?
	photos = filter(list)
	return photos
}

// some tools
func filter(slc []string) (res []string) {
	for _, v := range slc {
		if v != "" {
			res = append(res, v)
		}
	}
	return res
}
