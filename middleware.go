package main

import (
	"net/http"
	"test/Utils"
)

func AddHandleFunc(route string, content string, f func(http.ResponseWriter, *http.Request)) {
	pages = append(pages, page{route: route, content: content, f: f})
}

func URLHandler(w http.ResponseWriter, req *http.Request) {
	title := req.URL.Path
	for _, t := range pages {
		if t.route == title || t.route[1:] == title {
			if t.content == "" {
				t.content = "text/plain"
			}
			w.Header().Set("Content-Type", t.content)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			t.f(w, req)
			return
		}
	}
	Utils.Include("404.html", w, nil)
}

type page struct {
	route   string
	content string
	f       func(http.ResponseWriter, *http.Request)
}

var pages []page = []page{}
