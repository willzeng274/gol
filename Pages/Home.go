package Pages

import (
	"net/http"
	C "test/Components"
	"test/Utils"
)

func Home(w http.ResponseWriter, req *http.Request) {
	Utils.Include("home.html", w, []C.Component{})
}
