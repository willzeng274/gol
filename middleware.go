package main

import (
	"net/http"
	"test/Utils"
)

func AddHandleFunc(route string, content string, f func(http.ResponseWriter, *http.Request)) {
	pages[route] = page{route: route, content: content, f: f}
}

func URLHandler(w http.ResponseWriter, req *http.Request) {
	title := req.URL.Path
	if t, v := pages[title]; v {
		if t.content == "" {
			t.content = "text/plain"
		}
		w.Header().Set("Content-Type", t.content)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		t.f(w, req)
		return
	}
	Utils.Include("404.html", w, nil)
}

type page struct {
	route   string
	content string
	f       func(http.ResponseWriter, *http.Request)
}

var pages map[string]page = map[string]page{}
