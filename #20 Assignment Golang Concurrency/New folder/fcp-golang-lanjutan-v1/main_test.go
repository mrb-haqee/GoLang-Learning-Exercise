package main_test

import (
	main "a21hc3NpZ25tZW50"
	"a21hc3NpZ25tZW50/handler"
	"a21hc3NpZ25tZW50/model"
	"bufio"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const errTimeoutMessage = "Timeout: execute test take time more than 1 second"

var _ = Describe("API Student Portal", Ordered, func() {
	var server main.API

	BeforeEach(func() {
		server = main.NewAPI()
	})

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
				users = append(users, model.User{ID: data[0], Name: data[1], Role: data[2], StudyCode: data[3]})
			}
		}

		var newData string
		for _, u := range users {
			newData += u.ID + "_" + u.Name + "_" + u.Role + "_" + u.StudyCode + "\n"
		}

		if err = os.WriteFile(filepath.Join("data", "users.txt"), []byte(newData), 0644); err != nil {
			panic(err)
		}
	})

	Describe("/register", func() {
		When("send one of the required fields is empty", func() {
			It("should return a wrong body request", func() {
				cred := model.User{ID: "ATEST1111", Name: "test tama", Role: "", StudyCode: ""}
				register, _ := json.Marshal(cred)
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(register))
				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
				Expect(ErrorResponse.Error).To(Equal("ID, name, study code or role is empty"))
			})
		})

		When("send correct data to register", func() {
			It("should return a success message", func() {
				cred := model.User{ID: "ATEST1111", Name: "test tama", Role: "user", StudyCode: "MK"}
				register, _ := json.Marshal(cred)
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(register))
				server.Handler().ServeHTTP(w, r)

				SuccessResponse := model.SuccessResponse{}
				json.NewDecoder(w.Body).Decode(&SuccessResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
				Expect(SuccessResponse.Message).To(Equal("register success"))
			})
		})

		When("send duplicate data to register", func() {
			It("should return a duplicate data message", func() {
				cred := model.User{ID: "ATEST1111", Name: "test tama", Role: "user", StudyCode: "MK"}
				register, _ := json.Marshal(cred)
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(register))
				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
				Expect(ErrorResponse.Error).To(Equal("user id already exist"))
			})
		})
	})

	Describe("/login", func() {
		When("send id with data not exist in file", func() {
			It("should return a wrong body request", func() {
				cred := model.UserLogin{ID: "ATEST1122", Name: "test tama"}
				login, _ := json.Marshal(cred)
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
				Expect(ErrorResponse.Error).To(Equal("user not found"))
			})
		})

		When("send correct data to login", func() {
			It("should return a success message", func() {
				cred := model.User{ID: "ATEST1112", Name: "test tama", Role: "admin", StudyCode: "MK"}
				register, _ := json.Marshal(cred)
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(register))
				server.Handler().ServeHTTP(w, r)

				cred = model.User{ID: "ATEST1112", Name: "test tama"}

				login, _ := json.Marshal(cred)
				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(login))

				server.Handler().ServeHTTP(w, r)

				SuccessResponse := model.SuccessResponse{}
				json.NewDecoder(w.Body).Decode(&SuccessResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
				Expect(SuccessResponse.Message).To(Equal("login success"))
			})
		})
	})

	Describe("/logout", func() {
		When("user not login", func() {
			It("should return a error message", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/logout", nil)
				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusUnauthorized))
			})
		})

		When("user already login", func() {
			It("should success message logout", func() {
				credLogin := model.UserLogin{ID: "ATEST1111", Name: "test tama"}
				login, _ := json.Marshal(credLogin)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				cookieUserId := w.Result().Cookies()[0]
				cookieUserRole := w.Result().Cookies()[1]

				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodPost, "/logout", nil)

				r.AddCookie(cookieUserId)
				r.AddCookie(cookieUserRole)

				server.Handler().ServeHTTP(w, r)

				SuccessResponse := model.SuccessResponse{}
				json.NewDecoder(w.Body).Decode(&SuccessResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
				Expect(SuccessResponse.Message).To(Equal("logout success"))
			})
		})
	})

	Describe("/study-program", func() {
		When("user not login", func() {
			It("should return a error message", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodGet, "/study-program", nil)

				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusUnauthorized))
			})
		})

		When("user already login", func() {
			It("should return a success message", func() {
				credLogin := model.UserLogin{ID: "ATEST1112", Name: "test tama"}
				login, _ := json.Marshal(credLogin)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				cookieUserId := w.Result().Cookies()[0]
				cookieUserRole := w.Result().Cookies()[1]

				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodGet, "/study-program", nil)

				r.AddCookie(cookieUserId)
				r.AddCookie(cookieUserRole)

				server.Handler().ServeHTTP(w, r)

				var studyProgram []model.StudyData

				json.NewDecoder(w.Body).Decode(&studyProgram)
				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
				Expect(len(studyProgram)).To(Equal(15))
			})
		})
	})

	Describe("/user/add", func() {
		When("user not login", func() {
			It("should return a error message", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/user/add", nil)
				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusUnauthorized))
			})
		})

		When("user login with role not admin", func() {
			It("should return a error message", func() {
				credLogin := model.UserLogin{ID: "ATEST1111", Name: "test tama"}
				login, _ := json.Marshal(credLogin)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				cookieUserId := w.Result().Cookies()[0]
				cookieUserRole := w.Result().Cookies()[1]

				newUser := model.User{ID: "TESTTEST11", Name: "test jaya", Role: "user"}

				newUserJson, _ := json.Marshal(newUser)

				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodPost, "/user/add", bytes.NewReader(newUserJson))

				r.AddCookie(cookieUserId)
				r.AddCookie(cookieUserRole)

				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusUnauthorized))

			})
		})

		When("user login with role admin with incorrect list-study data", func() {
			It("should return a error message", func() {
				credLogin := model.UserLogin{ID: "ATEST1112", Name: "test tama"}
				login, _ := json.Marshal(credLogin)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				cookieUserId := w.Result().Cookies()[0]
				cookieUserRole := w.Result().Cookies()[1]

				credAdd := model.User{ID: "TEST11113", Name: "test tama", Role: "admin", StudyCode: "TEST"}
				add, _ := json.Marshal(credAdd)

				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodPost, "/user/add", bytes.NewReader(add))

				r.AddCookie(cookieUserId)
				r.AddCookie(cookieUserRole)

				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
				Expect(ErrorResponse.Error).To(Equal("study code not found"))
			})
		})

		When("user role admin add user with existing user ID", func() {
			It("should return a error message", func() {
				credLogin := model.UserLogin{ID: "ATEST1112", Name: "test tama"}
				login, _ := json.Marshal(credLogin)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				cookieUserId := w.Result().Cookies()[0]
				cookieUserRole := w.Result().Cookies()[1]

				credAdd := model.User{ID: "ATEST1111", Name: "test tama", Role: "admin", StudyCode: "TEST"}
				add, _ := json.Marshal(credAdd)

				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodPost, "/user/add", bytes.NewReader(add))

				r.AddCookie(cookieUserId)
				r.AddCookie(cookieUserRole)

				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
				Expect(ErrorResponse.Error).To(Equal("study code not found"))
			})
		})

		When("user role admin add user with correct data", func() {
			It("should return a success message", func() {
				credLogin := model.UserLogin{ID: "ATEST1112", Name: "test tama"}
				login, _ := json.Marshal(credLogin)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				cookieUserId := w.Result().Cookies()[0]
				cookieUserRole := w.Result().Cookies()[1]

				credAdd := model.User{ID: "TEST11130", Name: "test create user", Role: "user", StudyCode: "MK"}
				add, _ := json.Marshal(credAdd)

				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodPost, "/user/add", bytes.NewReader(add))

				r.AddCookie(cookieUserId)
				r.AddCookie(cookieUserRole)

				server.Handler().ServeHTTP(w, r)

				SuccessResponse := model.SuccessResponse{}
				json.NewDecoder(w.Body).Decode(&SuccessResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
				Expect(SuccessResponse.Message).To(Equal("add user success"))
			})
		})
	})

	Describe("/user/delete", func() {
		When("user not login", func() {
			It("should return a error message", func() {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodDelete, "/user/delete", nil)
				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusUnauthorized))
			})
		})

		When("user login with role not admin", func() {
			It("should return a error message", func() {
				credLogin := model.UserLogin{ID: "ATEST1112", Name: "test tama"}
				login, _ := json.Marshal(credLogin)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				cookieUserId := w.Result().Cookies()[0]
				cookieUserRole := w.Result().Cookies()[1]

				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodDelete, "/user/delete", nil)

				r.AddCookie(cookieUserId)
				r.AddCookie(cookieUserRole)

				server.Handler().ServeHTTP(w, r)

				ErrorResponse := model.ErrorResponse{}
				json.NewDecoder(w.Body).Decode(&ErrorResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
				Expect(ErrorResponse.Error).To(Equal("user id is empty"))
			})
		})

		When("user login with role admin with existing user id", func() {
			It("should return a success message", func() {
				credLogin := model.UserLogin{ID: "ATEST1112", Name: "test tama"}
				login, _ := json.Marshal(credLogin)

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(login))
				server.Handler().ServeHTTP(w, r)

				cookieUserId := w.Result().Cookies()[0]
				cookieUserRole := w.Result().Cookies()[1]

				w = httptest.NewRecorder()
				r = httptest.NewRequest(http.MethodDelete, "/user/delete?id=TEST11130", nil)

				r.AddCookie(cookieUserId)
				r.AddCookie(cookieUserRole)

				server.Handler().ServeHTTP(w, r)

				SuccessResponse := model.SuccessResponse{}
				json.NewDecoder(w.Body).Decode(&SuccessResponse)
				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
				Expect(SuccessResponse.Message).To(Equal("delete success"))
			})
		})
	})

	Describe("/get-weather", func() {
		ReportAfterEach(func(report SpecReport) {
			if report.RunTime > (1 * time.Second) {
				AbortSuite(errTimeoutMessage)
			}
		})

		When("the handler implement goroutine", func() {
			It("should execute function less than 1 second", func() {
				handler.GetWetherByRegionAPI = func(region string) (model.MainWeather, error) {
					time.Sleep(400 * time.Millisecond)
					return model.MainWeather{
						ID: 1,
						Weather: []model.Weather{{
							Main:        "Clouds",
							Description: "scattered clouds",
							Icon:        "03d",
						}},
					}, nil
				}

				start := time.Now()

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodGet, "/get-weather", nil)
				server.Handler().ServeHTTP(w, r)

				elapsed := time.Since(start)

				Expect(elapsed).To(BeNumerically("<", time.Second))

				var resp []model.MainWeather

				json.NewDecoder(w.Body).Decode(&resp)
				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
				Expect(resp).NotTo(BeNil())
				Expect(len(resp)).To(Equal(10))
			})
		})
	})
})
