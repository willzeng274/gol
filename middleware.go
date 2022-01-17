package main

import (
	"net/http"
	"test/Utils"
)

func AddHandleFunc(method string, route string, content string, f func(http.ResponseWriter, *http.Request)) {
	if method == "" {
		method = "GET"
	}
	pages[route] = page{route: route, content: content, f: f, method: method}
}

func URLHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	title := req.URL.Path
	if m, v := middlewares[title]; v {
		for _, fn := range m {
			called := false
			fn(w, req, func() {
				called = true
			})
			if called {
				return
			}
		}
	}
	if t, v := pages[title]; v && t.method == req.Method {
		if t.content == "" {
			t.content = "text/plain"
		}
		w.Header().Set("Content-Type", t.content)
		t.f(w, req)
		return
	}
	if req.Method == "GET" {
		Utils.Include("404.html", w, nil)
	} else {
		http.Error(w, "No", 404)
	}
}

type page struct {
	route   string
	content string
	f       func(http.ResponseWriter, *http.Request)
	method  string
}

var pages map[string]page = map[string]page{}

func AddMiddleware(route string, call func(http.ResponseWriter, *http.Request, func())) {
	middlewares[route] = append(middlewares[route], call)
}

var middlewares map[string][]func(http.ResponseWriter, *http.Request, func())
