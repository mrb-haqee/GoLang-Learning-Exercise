package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func TimeHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		w.Write([]byte(fmt.Sprintf("%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())))
	}) // TODO: replace this
}

func SayHelloHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		check:=strings.Contains(fmt.Sprint(r.URL),"/hello?")
		if check {
			w.Write([]byte(fmt.Sprintf("Hello, %s!", r.URL.Query()["name"][0])))
		} else {
			w.Write([]byte("Hello there"))
		}
	}) // TODO: replace this
}

func main() {
	// TODO: answer here
	http.HandleFunc("/time", TimeHandler())
	http.HandleFunc("/hello", SayHelloHandler())
	http.ListenAndServe("localhost:8080", nil)
}
