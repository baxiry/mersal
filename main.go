package main

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func main() {

	db = ConnectDB()
	defer db.Close()

	e := echo.New()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.Renderer = Templs("templates")

	// files
	e.Static("/assets", Assets())
	e.Static("/fs", PhotoFold())

	// account and verefy
	e.GET("/", HomePage)
	e.GET("/sign", SignPage)
	e.POST("/sign", Signup)
	e.GET("/login", LoginPage)
	e.POST("/login", Login)
	e.GET("/user/:id", Profile)
	e.GET("/fotos", PhotosPage)
	e.POST("/upfotos/:id", UpPhotos)
	e.GET("/upacount", UpdatePage)
	e.POST("/upacount", Update)

	e.GET("/messages", MessagesPage)
	e.GET("/activity", ActivityPage)
	e.GET("/search", SearchPage)

	//e.POST("/updatefotos/:id", updateProdFotos)

	//e.GET("/:catigory/:id", getOneProd) // whech is beter ? :catigory or /product ?

	e.Logger.Fatal(e.Start(":8080"))

}
