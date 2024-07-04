package gee

import (
	"fmt"
	"net/http"
)

func InitLoginInterface(r *Engine) {

	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)

		fmt.Print("登陆")
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {

			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
}
