package main

import (
	"fmt"
	"net/http"
	"os"
)

func MethodGet(r *http.Request) error {
	if r.Method != http.MethodGet {
		return fmt.Errorf("Method not allowed")
	}
	return nil
}

func CheckDataRequest(r *http.Request) error {
	data := r.URL.Query().Get("data")
	if len(data) == 0 {
		return fmt.Errorf("Data not found")
	}
	return nil
}

func CheckOpenFile(r *http.Request) error {
	filename := r.URL.Query().Get("filename")
	_, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("File not found")
	}
	return nil
}

func MethodHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := MethodGet(r)
		if err != nil {
			w.WriteHeader(405)
			fmt.Fprint(w,err)
		}else {
			w.WriteHeader(200)
			fmt.Fprint(w, "Method handler passed")
		}
	}
}

func DataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := CheckDataRequest(r)
		if err != nil {
			w.WriteHeader(404)
			fmt.Fprint(w,err)
		}else {
			w.WriteHeader(200)
			fmt.Fprint(w, "Data handler passed")
		}
	}
}

func OpenFileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := CheckOpenFile(r)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w,err)
		}else {
			w.WriteHeader(200)
			fmt.Fprint(w, "Error handler passed")
		}
	}
}
