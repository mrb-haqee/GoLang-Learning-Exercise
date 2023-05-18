package main

import (
	"fmt"
	"net/http"
	"strings"
)

var students = []string{
	"Aditira",
	"Dito",
	"Afis",
	"Eddy",
}

func IsNameExists(name string) bool {
	for _, n := range students {
		check:=strings.Contains(name, n)
		if check{
			return check
		}
	}
	return false
}

func CheckStudentName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		check := IsNameExists(fmt.Sprint(r.URL))
		switch r.Method {
		case "GET":
			if check{
				w.WriteHeader(200)
				w.Write([]byte("Name is exists"))
			} else{
				w.WriteHeader(404)
				w.Write([]byte("Data not found"))
			}
		case "POST":
			w.WriteHeader(405)
			w.Write([]byte("Method is not allowed"))
		}
	} // TODO: replace this
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/students", CheckStudentName())
	// TODO: answer here
	return mux
}

func main() {
	fmt.Println("Succses")

	http.ListenAndServe("localhost:8080", GetMux())
}
