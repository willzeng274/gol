package Utils

import (
	"fmt"
	"net/http"
	"os"
	C "test/Components"
	P "test/parsers"
)

func Include(file string, w http.ResponseWriter, Components []C.Component) {
	html, err := os.ReadFile("templates/" + file)
	Handler(err)
	raw := string(html)
	for _, c := range Components {
		P.Parse(&raw, c)
	}
	fmt.Fprintf(w, raw)
}
