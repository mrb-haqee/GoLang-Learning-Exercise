package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		Weekday := time.Now().Weekday()
		Day := time.Now().Day()
		Month := time.Now().Month()
		Year := time.Now().Year()
		writer.Write([]byte(fmt.Sprint(Weekday)+", "+fmt.Sprint(Day)+" "+fmt.Sprint(Month)+" "+fmt.Sprint(Year)))
	} // TODO: replace this
}

func main() {
	http.ListenAndServe("localhost:8080", GetHandler())

}
