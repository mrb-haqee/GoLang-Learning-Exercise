package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type QuotesHandler struct {
	Quotes []string
}

// type

var Quotes = []string{
	"Be yourself; everyone else is already taken. ― Oscar Wilde",
	"Be the change that you wish to see in the world. ― Mahatma Gandhi",
	"I have not failed. I've just found 10,000 ways that won't work. ― Thomas A. Edison",
	"It is never too late to be what you might have been. ― George Eliot",
	"Everything you can imagine is real. ― Pablo Picasso",
	"Nothing is impossible, the word itself says 'I'm possible'! ― Audrey Hepburn",
}

// TODO: answer here

func (qh QuotesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	quote := Quotes[rand.Intn(len(Quotes))]
	w.Write([]byte(quote))
	// TODO: answer here
}

func main() {
	// refrensi : https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go
	mux := http.NewServeMux()
	handler := QuotesHandler{}
	mux.Handle("/", handler)

	fmt.Println("succses")
	http.ListenAndServe("localhost:8080", mux)
}
