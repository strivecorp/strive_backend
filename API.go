package main

import (
	"log"
	"net/http"
)

func APIHandler() {
	http.ListenAndServe(":8080", nil)
}