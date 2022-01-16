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

// package main

// import (
// 	"fmt"
// 	t "test/testmod"
// )

// func main() {
// 	fmt.Println("Hello World")
// 	fmt.Println(Hello)
// 	t.Ok()
// 	a := 1
// 	b := 2
// 	c1 := make(chan int)
// 	go func() {
// 		res := Add(&a, &b)
// 		c1 <- res
// 	}()
// 	select {
// 	case msg := <-c1:
// 		fmt.Println(msg)
// 	}
// }
