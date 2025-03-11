package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/johnnyfreeman/learning-go-api-server/templates"
)

func main() {
	component := templates.Hello("John")

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
