package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func TimeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_=r
		w.Write([]byte(fmt.Sprintf("%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())))	
	} // TODO: replace this
}

func SayHelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		check:=strings.Contains(fmt.Sprint(r.URL),"/hello?")
		if check {
			w.Write([]byte(fmt.Sprintf("Hello, %s!", r.URL.Query()["name"][0])))
		} else {
			w.Write([]byte("Hello there"))
		}

	} // TODO: replace this
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/time", TimeHandler())
	mux.HandleFunc("/hello", SayHelloHandler())
	// TODO: answer here
	return mux
}

func main() {
	fmt.Println("succses")
	http.ListenAndServe("localhost:8080", GetMux())
}
