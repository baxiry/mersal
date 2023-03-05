package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
)

var (
	AppName   = "social"
	TableName = "users"
	db        *sql.DB
)

// init database
func ConnectDB() *sql.DB {
	var err error

	db, err = sql.Open(
		"mysql", "root:123456@tcp(127.0.0.1:3306)/?charset=utf8&parseTime=True&loc=Local")
	if err != nil { // why no error when db is not runinig ??
		fmt.Println("run mysql server", err)
		// TODO report this error.

		// wehen db is stoped no error is return.
		// we expecte errore no database is runing

		// my be this error is fixed with panic ping pong bellow
	}

	if err = db.Ping(); err != nil {
		// TODO handle this error: dial tcp 127.0.0.1:3306: connect: connection refused
		fmt.Println("mybe database is not runing or error is: ", err)
		os.Exit(1)
	}
	return db
}

// init templates

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// path file is depends to enveronment.

func Templs(path string) *Template {
	return &Template{templates: template.Must(template.ParseFiles(listFiles(path)...))}
}

// listFiles return list filenames os spicific dir
// use paht.wolkFile insteade

func listFiles(dir string) (list []string) {

	f, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	sublist := make([]string, 0)
	root := dir + "/"
	for _, v := range files {
		root = dir + "/"
		if v.IsDir() {
			root = root + v.Name()
			sublist = listFiles(root)
			//for _, filename := range sublist {
			//	list = append(list, filename)
			//}
			continue
		}
		list = append(list, root+v.Name())
	}
	for _, f := range sublist {
		list = append(list, f)
	}

	return list
}

// folder when photos is stored.

func PhotoFold() string {
	//if os.Getenv("USERNAME") == "fedor" {
	//	return "/home/fedor/repo/files/"
	//}
	return "../files/"
}

// where is assets  path ?
func Assets() string {
	home, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	if home != "/Users/fedora/repo/social" {
		return "/root/social/assets"
	}
	return "assets"
}

// CREATE DATABASE (not with sqlite);
func CreateDB(dbName string, db *sql.DB) {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName + ";")
	if err != nil {
		panic(err)
	}
}
