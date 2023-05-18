package main

import (
	"fmt"
	"net/http"
)

func StudentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Student page"))
	}
}

func RequestMethodGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case "POST":
            w. WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method is not allowed"))
		case "PUT":
            w. WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method is not allowed"))
		case "DELETE":
            w. WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method is not allowed"))
		case "GET":
			next.ServeHTTP(w,r)
		}
	}) // TODO: replace this
}

func main() {
	// TODO: answer here
	//refrensi: https://www.alexedwards.net/blog/making-and-using-middleware
	rev:=RequestMethodGet(StudentHandler()).ServeHTTP
	http.HandleFunc("/student", rev)

	fmt.Println("succses")
	http.ListenAndServe("localhost:8080", nil)
}
