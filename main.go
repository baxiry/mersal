package main

import (
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Static file server
	app.Static("/", "./files")

	log.Fatal(app.Listen(3000))
}

//"net/http"
// Simple static webserver:
//log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/home/fedora/Videos"))))
