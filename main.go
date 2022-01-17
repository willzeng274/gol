package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	P "test/Pages"
)

func main() {
	http.HandleFunc("/", URLHandler)
	AddHandleFunc("GET", "/", "text/html", P.Home)
	AddHandleFunc("GET", "/website", "text/html", P.Website)
	var PORT string
	if _, err := os.Stat(".env"); errors.Is(err, os.ErrNotExist) {
		fmt.Println("Enter PORT (default is 8090)")
		fmt.Scanln(&PORT)
	} else {
		envVals, err := os.ReadFile(".env")
		if err != nil {
			panic(err)
		}
		ParseEnv(string(envVals))
		PORT = os.Getenv("PORT")
	}
	if PORT == "" {
		PORT = "8090"
	}
	fmt.Printf("Running server on localhost:%s", PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil)
}
