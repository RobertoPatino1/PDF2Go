package main

import (
	"log"
	"net/http"

	"github.com/RobertoPatino1/md_pdf_2_Go/core/handler"
)

func main() {
    http.HandleFunc("/convert", handler.ConvertHandler)
    log.Println("Server listening on port:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
