package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
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

	// chat app

	fmt.Println("version 0.0.2\nim start at :8080")

	http.HandleFunc("/ws", wsHandler)

	go panic(http.ListenAndServe(":8080", nil))

	e.Logger.Fatal(e.Start(":8080"))

}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	//if r.Header.Get("Origin")!="http://"+r.Host {http.Error(w,"Origin not allowed",-1);return}

	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", 404)
	}
	go im.ServeMessages(conn)
}
