package main_test

import (
	main "a21hc3NpZ25tZW50"
	"a21hc3NpZ25tZW50/model"
	"bufio"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("API Student Portal", Ordered, func() {
	AfterAll(func() {
		f, err := os.Open(filepath.Join("data", "users.txt"))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		var users []model.User

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			data := strings.Split(scanner.Text(), "_")
			if data[0] != "ATEST1111" && data[0] != "ATEST1112" && data[0] != "TEST11130" {
				users = append(users, model.User{ID: data[0], Name: data[1], StudyCode: data[3]})
			}
		}

		var newData string
		for _, u := range users {
			newData += u.ID + "_" + u.Name + "_" + u.StudyCode + "\n"
		}

		if err = os.WriteFile(filepath.Join("data", "users.txt"), []byte(newData), 0644); err != nil {
			panic(err)
		}
	})

	Describe("/study-program", func() {
		When("method is not GET", func() {
			It("should return a error message", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/study-program", nil)

				handler := http.HandlerFunc(main.GetStudyProgram())
				handler.ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusMethodNotAllowed))
				Expect(ErrorResponse.Error).To(Equal("Method is not allowed!"))
			})
		})

		When("method is GET", func() {
			It("should return a success message", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodGet, "/study-program", nil)

				handler := http.HandlerFunc(main.GetStudyProgram())
				handler.ServeHTTP(w, r)

				var studyProgram []model.StudyData

				json.NewDecoder(w.Body).Decode(&studyProgram)
				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
				Expect(len(studyProgram)).To(Equal(15))
			})
		})
	})

	Describe("/user/add", func() {
		When("method is not POST", func() {
			It("should return a error message", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodGet, "/user/add", nil)

				handler := http.HandlerFunc(main.AddUser())
				handler.ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusMethodNotAllowed))
				Expect(ErrorResponse.Error).To(Equal("Method is not allowed!"))
			})
		})

		When("method is POST with send empty body request", func() {
			It("should return a error message", func() {
				newUser := model.User{}

				newUserJson, _ := json.Marshal(newUser)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/user/add", bytes.NewReader(newUserJson))

				handler := http.HandlerFunc(main.AddUser())
				handler.ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))

			})
		})

		When("method is POST with incorrect list-study data", func() {
			It("should return a error message", func() {
				credAdd := model.User{ID: "TEST11113", Name: "test tama", StudyCode: "TEST"}
				add, _ := json.Marshal(credAdd)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/user/add", bytes.NewReader(add))

				handler := http.HandlerFunc(main.AddUser())
				handler.ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
				Expect(ErrorResponse.Error).To(Equal("study code not found"))
			})
		})

		When("method is POST with existing user ID", func() {
			It("should return a error message", func() {
				credAdd := model.User{ID: "ATEST1111", Name: "test tama", StudyCode: "TEST"}
				add, _ := json.Marshal(credAdd)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/user/add", bytes.NewReader(add))

				handler := http.HandlerFunc(main.AddUser())
				handler.ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
				Expect(ErrorResponse.Error).To(Equal("study code not found"))
			})
		})

		When("method is POST with correct data", func() {
			It("should return a success message", func() {
				credAdd := model.User{ID: "TEST11130", Name: "test create user", StudyCode: "MK"}
				add, _ := json.Marshal(credAdd)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/user/add", bytes.NewReader(add))

				handler := http.HandlerFunc(main.AddUser())
				handler.ServeHTTP(w, r)

				SuccessResponse := model.SuccessResponse{}
				json.NewDecoder(w.Body).Decode(&SuccessResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
				Expect(SuccessResponse.Message).To(Equal("add user success"))
			})
		})
	})

	Describe("/user/delete", func() {
		When("method is not DELETE", func() {
			It("should return a error message", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/user/delete", nil)

				handler := http.HandlerFunc(main.DeleteUser())
				handler.ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusMethodNotAllowed))
				Expect(ErrorResponse.Error).To(Equal("Method is not allowed!"))
			})
		})

		When("method is DELETE with send empty query param", func() {
			It("should return a error message", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodDelete, "/user/delete", nil)

				handler := http.HandlerFunc(main.DeleteUser())
				handler.ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
				Expect(ErrorResponse.Error).To(Equal("user id is empty"))
			})
		})

		When("method is DELETE with not existing user id", func() {
			It("should return a error message", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodDelete, "/user/delete?id=A9999", nil)

				handler := http.HandlerFunc(main.DeleteUser())
				handler.ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
				Expect(ErrorResponse.Error).To(Equal("user id not found"))
			})
		})

		When("method is DELETE with existing user id", func() {
			It("should return a success message", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodDelete, "/user/delete?id=TEST11130", nil)

				handler := http.HandlerFunc(main.DeleteUser())
				handler.ServeHTTP(w, r)

				SuccessResponse := model.SuccessResponse{}
				json.NewDecoder(w.Body).Decode(&SuccessResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
				Expect(SuccessResponse.Message).To(Equal("delete success"))
			})
		})
	})
})
