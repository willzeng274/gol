package main

import (
	"fmt"
	"net/http"
	P "test/Pages"
)

func main() {
	fmt.Println("Running server on localhost:8090")

	http.HandleFunc("/", URLHandler)
	AddHandleFunc("/", "text/html", P.Home)
	AddHandleFunc("/website", "text/html", P.Website)

	http.ListenAndServe(":8090", nil)
}
