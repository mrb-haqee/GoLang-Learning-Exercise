package main

import (
	"a21hc3NpZ25tZW50/handler"
	"a21hc3NpZ25tZW50/middleware"
	"fmt"
	"net/http"
)

type API struct {
	mux *http.ServeMux
}

func NewAPI() API {
	mux := http.NewServeMux()
	api := API{
		mux,
	}

	mux.Handle("/register", middleware.Post(http.HandlerFunc(handler.Register)))
	mux.Handle("/login", middleware.Post(http.HandlerFunc(handler.Login)))
	mux.Handle("/logout", middleware.Post(handler.Auth(http.HandlerFunc(handler.Logout))))

	mux.Handle("/study-program", middleware.Get(handler.Auth(http.HandlerFunc(handler.GetStudyProgram))))
	mux.Handle("/user/add", middleware.Post(handler.Auth(handler.AuthAdmin(http.HandlerFunc(handler.AddUser)))))
	mux.Handle("/user/delete", middleware.Delete(handler.Auth(handler.AuthAdmin(http.HandlerFunc(handler.DeleteUser)))))

	mux.Handle("/get-weather", middleware.Get(http.HandlerFunc(handler.GetWeather)))
	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}

func main() {
	mainAPI := NewAPI()
	mainAPI.Start()
}
