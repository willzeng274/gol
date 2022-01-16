package Pages

import (
	"net/http"
	C "test/Components"
	"test/Utils"
)

func Website(w http.ResponseWriter, req *http.Request) {
	Utils.Include("index.html", w, []C.Component{C.Navbar})
}
