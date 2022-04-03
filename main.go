package main

import (
	"net/http"

	"github.com/alexandrevicenzi/go-sse"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
)

var (
	ssevent = sse.NewServer(&sse.Options{}) // pass nil to show default Options
	data    string
)

func main() {
	db := setdb()
	defer db.Close()
	defer ssevent.Shutdown()

	e := echo.New()
	// middleware
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	go func() {
		for {
			data = <-dataPipe
			channel = gjson.Get(data, "channel").String()
			msg := gjson.Get(data, "msg").String()
			ssevent.SendMessage("/events/"+channel, sse.SimpleMessage(msg))
		}
	}()

	// routers
	e.Any("/events/:channel", pusher)
	e.POST("/", messages)
	e.GET("/", home)
	e.POST("/signup", signup)
	e.POST("/login", login)

	e.Logger.Fatal(e.Start(":3000"))
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
