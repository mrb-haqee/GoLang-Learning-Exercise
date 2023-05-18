package main

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func GetStudyProgram() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			read, err := os.ReadFile("data/list-study.txt")
			if err != nil {
				panic(err)
			}
			readArr := strings.Split(string(read), "\n")
			DataArr := []model.StudyData{}
			for i, v := range readArr {
				DataSplit := strings.Split(v, "_")
				if i == len(readArr)-1 {
					DataArr = append(DataArr, model.StudyData{Code: DataSplit[0], Name: string(DataSplit[1][:len(DataSplit[1])])})
					break
				}
				DataArr = append(DataArr, model.StudyData{Code: DataSplit[0], Name: string(DataSplit[1][:len(DataSplit[1])-1])})
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err = json.NewEncoder(w).Encode(DataArr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		default:
			MesErr := model.ErrorResponse{Error: "Method is not allowed!"}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			err := json.NewEncoder(w).Encode(MesErr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

		}
		// TODO: answer here
	}
}

func AddUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				w.Write([]byte("Error read Body" + err.Error()))
			}
			var user model.User
			err = json.Unmarshal(body, &user)
			if err != nil {
				panic(err)
			}
			dataUser, err := os.ReadFile("data/users.txt")
			if err != nil {
				panic(err)
			}
			dataStudy, err := os.ReadFile("data/list-study.txt")
			if err != nil {
				panic(err)
			}
			dataStudyArr := strings.Split(string(dataStudy), "\n")
			checkStudy := false
			for _, v := range dataStudyArr {
				DataSplit := strings.Split(v, "_")
				if DataSplit[0] == user.StudyCode {
					checkStudy = true
					break
				}
			}
			if user.ID == "" || user.Name == "" || user.StudyCode == "" {
				MesErr := model.ErrorResponse{Error: "ID, name, or study code is empty"}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(400)
				err := json.NewEncoder(w).Encode(MesErr)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			} else if strings.Contains(string(dataUser), user.ID) {
				MesErr := model.ErrorResponse{Error: "user id already exist"}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(400)
				err := json.NewEncoder(w).Encode(MesErr)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			} else if !checkStudy {
				MesErr := model.ErrorResponse{Error: "study code not found"}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(400)
				err := json.NewEncoder(w).Encode(MesErr)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			} else {
				input := fmt.Sprint(user.ID + "_" + user.Name + "_" + user.StudyCode)
				err := os.WriteFile("data/users.txt", []byte(input), 0644)
				if err != nil {
					panic(err)
				}
				MesErr := model.SuccessResponse{Username: user.Name, Message: "add user success"}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				err = json.NewEncoder(w).Encode(MesErr)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		default:
			MesErr := model.ErrorResponse{Error: "Method is not allowed!"}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			err := json.NewEncoder(w).Encode(MesErr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
		// TODO: answer here
	}
}

func DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "DELETE":
			DeleteID := r.URL.Query().Get("id")
			readUser, err := os.ReadFile("data/users.txt")
			if err != nil {
				panic(err)
			}
			dataUserArr := strings.Split(string(readUser), "\n")
			CheckUser := false
			nameUser := ""
			for _, v := range dataUserArr {
				if strings.Contains(v, DeleteID) {
					nameUser, _, _ = strings.Cut(v, "_")
					CheckUser = true
				}
			}
			if DeleteID == "" {
				MesErr := model.ErrorResponse{Error: "user id is empty"}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(400)
				err := json.NewEncoder(w).Encode(MesErr)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			} else if !CheckUser {
				MesErr := model.ErrorResponse{Error: "user id not found"}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(400)
				err := json.NewEncoder(w).Encode(MesErr)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			} else {
				err := os.Remove("data/users.txt")
				if err != nil {
					panic(err)
				}
				file, err := os.Create("data/users.txt")
				if err != nil {
					panic(err)
				}
				defer file.Close()

				for _, v := range dataUserArr {
					if strings.Contains(v, nameUser) {
						continue
					}
					_, err = file.WriteString(v)
					if err != nil {
						panic(err)
					}
				}

				MesErr := model.SuccessResponse{Username: nameUser, Message: "delete success"}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(200)
				err = json.NewEncoder(w).Encode(MesErr)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		default:
			MesErr := model.ErrorResponse{Error: "Method is not allowed!"}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			err := json.NewEncoder(w).Encode(MesErr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
		// TODO: answer here
	}
}

func main() {
	http.HandleFunc("/study-program", GetStudyProgram())
	// http.HandleFunc("/user/add", AddUser())
	// http.HandleFunc("/user/delete", DeleteUser())

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
