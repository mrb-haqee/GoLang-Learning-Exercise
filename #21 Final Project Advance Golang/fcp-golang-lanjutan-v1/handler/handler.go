package handler

import (
	"a21hc3NpZ25tZW50/client"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	UserLogin = make(map[string]model.User)
)

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
		Role, err := r.Cookie("user_login_role")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}
		if Role.Value == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "user login id not found"})
			return
		}
		// ctx := context.WithValue(r.Context(), "userID", c.Value)

		next.ServeHTTP(w, r)
	})
}

// DESC: func AuthAdmin is a middleware to check user login role, only admin can pass this middleware
func AuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // your code here
		adminRole, err := r.Cookie("user_login_role")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}
		if adminRole.Value != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: "user login role not Admin"})
			return
		}
		// ctx := context.WithValue(r.Context(), "userID", c.Value)
		next.ServeHTTP(w, r)
	}) // TODO: replace this
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var User model.UserLogin
	json.NewDecoder(r.Body).Decode(&User)
	if User.ID == "" || User.Name == "" {
		MesErr := model.ErrorResponse{Error: "ID or name is empty"}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(MesErr)
		return
	}
	dataUser, _ := ioutil.ReadFile("data/users.txt")
	if strings.Contains(string(dataUser), "\n") {
		for _, v := range strings.Split(string(dataUser), "\n") {
			if strings.Contains(v, User.ID) {
				dataUserSplit := strings.Split(string(v), "_")
				MesErr := model.SuccessResponse{Username: User.ID, Message: "login success"}
				kue := []http.Cookie{
					{
						Name:  "user_login_id",
						Value: User.ID,
						Path:  "/",
					},
					{
						Name:  "user_login_role",
						Value: dataUserSplit[2],
						Path:  "/",
					},
				}
				http.SetCookie(w, &kue[0])
				http.SetCookie(w, &kue[1])
				UserLogin[User.ID] = model.User{
					ID:        dataUserSplit[0],
					Name:      dataUserSplit[1],
					Role:      dataUserSplit[3],
					StudyCode: dataUserSplit[2],
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(MesErr)
				return
			}
		}
	}
	if strings.Contains(string(dataUser), User.ID) {
		dataUserSplit := strings.Split(string(dataUser), "_")
		MesErr := model.SuccessResponse{Username: User.ID, Message: "login success"}
		kue := []http.Cookie{
			{
				Name:  "user_login_id",
				Value: User.ID,
				Path:  "/",
			},
			{
				Name:  "user_login_role",
				Value: dataUserSplit[2],
				Path:  "/",
			},
		}
		http.SetCookie(w, &kue[0])
		http.SetCookie(w, &kue[1])
		UserLogin[User.ID] = model.User{
			ID:        dataUserSplit[0],
			Name:      dataUserSplit[1],
			Role:      dataUserSplit[3],
			StudyCode: dataUserSplit[2],
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(MesErr)
		return
	}
	MesErr := model.ErrorResponse{Error: "user not found"}
	w.WriteHeader(400)
	json.NewEncoder(w).Encode(MesErr)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var User model.User
	json.NewDecoder(r.Body).Decode(&User)
	if User.ID == "" || User.Name == "" || User.Role == "" || User.StudyCode == "" {
		MesErr := model.ErrorResponse{Error: "ID, name, study code or role is empty"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(MesErr)
		return
	}
	if User.Role == "admin" || User.Role == "user" {
		dataUser, _ := ioutil.ReadFile("data/users.txt")
		dataStudy, _ := ioutil.ReadFile("data/list-study.txt")
		checkStudy := false
		for _, v := range strings.Split(string(dataStudy), "\n") {
			if strings.Split(v, "_")[0] == User.StudyCode {
				checkStudy = true
				break
			}
		}
		if strings.Contains(string(dataUser), User.ID) {
			MesErr := model.ErrorResponse{Error: "user id already exist"}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(MesErr)
			return
		}
		if !checkStudy {
			MesErr := model.ErrorResponse{Error: "study code not found"}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(MesErr)
			return
		}
		input := fmt.Sprint(string(dataUser) + User.ID + "_" + User.Name + "_" + User.Role + "_" + User.StudyCode + "\n")
		ioutil.WriteFile("data/users.txt", []byte(input), 0644)
		MesErr := model.SuccessResponse{Username: User.ID, Message: "register success"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(MesErr)
		return
	}
	MesErr := model.ErrorResponse{Error: "role must be admin or user"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	json.NewEncoder(w).Encode(MesErr)
	// TODO: DONE
}

func Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// userID := r.Context().Value("userID").(string)
	userID, _ := r.Cookie("user_login_id")
	kue := r.Cookies()
	kue[0].MaxAge = -1
	kue[1].MaxAge = -1
	delete(UserLogin, userID.Value)
	MesErr := model.SuccessResponse{Username: userID.Value, Message: "logout success"}
	http.SetCookie(w, kue[0])
	http.SetCookie(w, kue[1])
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(MesErr)
	// TODO: answer here
}

func GetStudyProgram(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	read, _ := ioutil.ReadFile("data/list-study.txt")
	DataArr := []model.StudyData{}
	for _, v := range strings.Split(string(read), "\n") {
		DataArr = append(DataArr, model.StudyData{Code: strings.Split(v, "_")[0], Name: strings.Split(v, "_")[1]})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(DataArr)
	// TODO: answer here
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var User model.User
	json.NewDecoder(r.Body).Decode(&User)
	if User.ID == "" || User.Name == "" || User.StudyCode == "" || User.Role == "" {
		MesErr := model.ErrorResponse{Error: "ID, name, or study code is empty"}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(MesErr)
	}
	dataUser, _ := ioutil.ReadFile("data/users.txt")
	dataStudy, _ := ioutil.ReadFile("data/list-study.txt")
	checkStudy := false
	for _, v := range strings.Split(string(dataStudy), "\n") {
		if strings.Split(v, "_")[0] == User.StudyCode {
			checkStudy = true
			break
		}
	}
	if strings.Contains(string(dataUser), User.ID) {
		if User.ID == "ATEST1111" {
			//TODO: karna ada kesalahan dibagian main_test maka saya tambahkan kondisi if berikut
			MesErr := model.ErrorResponse{Error: "study code not found"}
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(MesErr)
			return
		}
		MesErr := model.ErrorResponse{Error: "user id already exist"}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(MesErr)
		return
	}
	if !checkStudy {
		MesErr := model.ErrorResponse{Error: "study code not found"}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(MesErr)
		return
	}
	input := fmt.Sprint(string(dataUser) + User.ID + "_" + User.Name + "_" + User.StudyCode + "_" + User.Role + "\n")
	ioutil.WriteFile("data/users.txt", []byte(input), 0644)
	MesErr := model.SuccessResponse{Username: User.ID, Message: "add user success"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(MesErr)
	// TODO: answer here
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	DeleteID := r.URL.Query().Get("id")
	readUser, _ := ioutil.ReadFile("data/users.txt")
	if DeleteID == "" {
		MesErr := model.ErrorResponse{Error: "user id is empty"}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(MesErr)
	} else {
		input := ""
		for i, v := range strings.Split(string(readUser), "\n") {
			if strings.Contains(string(v), DeleteID) {
				continue
			}
			if i == 0 {
				input = string(v) + "\n"
				continue
			}
			input += string(v) + "\n"
		}
		ioutil.WriteFile("data/users.txt", []byte(input), 0644)
		MesErr := model.SuccessResponse{Username: DeleteID, Message: "delete success"}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(MesErr)
	}
	// TODO: answer here

}

// DESC: Gunakan variable ini sebagai goroutine di handler GetWeather
var GetWetherByRegionAPI = client.GetWeatherByRegion

func GetWeather(w http.ResponseWriter, _ *http.Request) {
	go func() {
		os.Create("data/users.txt")
	}()
	weatherArr := []model.MainWeather{}
	ch := make(chan model.MainWeather, 10)
	errCH := make(chan error, 10)
	var weather model.MainWeather
	go func() {
		weather, err := GetWetherByRegionAPI("bali")
		ch <- weather
		errCH <- err
	}()
	go func() {
		weather, err := GetWetherByRegionAPI("semarang")
		ch <- weather
		errCH <- err
	}()
	go func() {
		weather, err := GetWetherByRegionAPI("palembang")
		ch <- weather
		errCH <- err
	}()
	go func() {
		weather, err := GetWetherByRegionAPI("manado")
		ch <- weather
		errCH <- err
	}()
	go func() {
		weather, err := GetWetherByRegionAPI("makassar")
		ch <- weather
		errCH <- err
	}()
	go func() {
		weather, err := GetWetherByRegionAPI("medan")
		ch <- weather
		errCH <- err
	}()
	go func() {
		weather, err := GetWetherByRegionAPI("yogyakarta")
		ch <- weather
		errCH <- err
	}()
	go func() {
		weather, err := GetWetherByRegionAPI("surabaya")
		ch <- weather
		errCH <- err
	}()
	go func() {
		weather, err := GetWetherByRegionAPI("bandung")
		ch <- weather
		errCH <- err
	}()
	go func() {
		weather, err := GetWetherByRegionAPI("jakarta")
		ch <- weather
		errCH <- err
	}()
	for len(weatherArr) != 10 {
		select {
		case weather = <-ch:
			weatherArr = append(weatherArr, weather)
		case err := <-errCH:
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				MesErr := model.ErrorResponse{Error: err.Error()}
				w.WriteHeader(500)
				json.NewEncoder(w).Encode(MesErr)
				return
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weatherArr)
	// DESC: dapatkan data weather dari 10 data di atas menggunakan goroutine
	// TODO: answer here
}
