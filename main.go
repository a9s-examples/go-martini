package main

import (
	"github.com/go-martini/martini"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = "8080"
	}

	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello, World!"
	})
	m.RunOnAddr(":" + port)
}
