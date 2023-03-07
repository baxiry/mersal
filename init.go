package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

const (
	AppName = "social"
	DBName  = "mersal"
	//TableName = "users"
)

var (
	db *sql.DB
)

// init database
func ConnectDB() *sql.DB {
	// sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/?charset=utf8&parseTime=True&loc=Local")
	db, err := sql.Open("sqlite3", "mersal.db")

	if err != nil { // why no error when db is not runinig ??
		fmt.Println("run mysql server", err)

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

	if home != "/Users/fedora/repo/mersal" {
		return "/root/mersal/assets"
	}
	return "assets"
}

// CREATE DATABASE (not with sqlite);
func CreateDB(dbName string, db *sql.DB) {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + DBName + ";")
	if err != nil {
		panic(err)
	}
}

// CREATE DATABASE (not with sqlite);
func CreateTable(tableName string, db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + tableName + ";")
	if err != nil {
		panic(err)
	}
}
