package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	dataPipe = make(chan string, 1)
	channel  string
	msg      string
)

func pusher(c echo.Context) error {
	ssevent.ServeHTTP(c.Response(), c.Request())
	return nil
}

func messages(c echo.Context) error {
	dataPipe <- c.FormValue("data")

	return c.String(http.StatusOK, "ok")
}
