package handler

import (
	"a21hc3NpZ25tZW50/client"
	"a21hc3NpZ25tZW50/model"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var UserLogin = make(map[string]model.User)

// DESC: func Auth is a middleware to check user login id, only user that already login can pass this middleware
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("user_login_id")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}

		if _, ok := UserLogin[c.Value]; !ok || c.Value == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "user login id not found"})
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "userID", c.Value)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// DESC: func AuthAdmin is a middleware to check user login role, only admin can pass this middleware
func AuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // your code here }) // TODO: replace this
}

func Login(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
}

func Register(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
}

func Logout(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)

	// TODO: answer here
}

func GetStudyProgram(w http.ResponseWriter, r *http.Request) {
	// list study program
	// TODO: answer here
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
}

// DESC: Gunakan variable ini sebagai goroutine di handler GetWeather
var GetWetherByRegionAPI = client.GetWeatherByRegion

func GetWeather(w http.ResponseWriter, r *http.Request) {
	var listRegion = []string{"jakarta", "bandung", "surabaya", "yogyakarta", "medan", "makassar", "manado", "palembang", "semarang", "bali"}

	// DESC: dapatkan data weather dari 10 data di atas menggunakan goroutine
	// TODO: answer here
}




