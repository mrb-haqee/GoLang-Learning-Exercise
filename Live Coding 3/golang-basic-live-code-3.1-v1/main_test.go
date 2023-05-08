package main_test

import (
	main "a21hc3NpZ25tZW50"
	"a21hc3NpZ25tZW50/model"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	var (
		user1 model.User
		user2 model.User
	)
	app := main.NewLearnly([]model.User{}, []model.Lesson{})

	BeforeEach(func() {
		user1 = model.User{
			Name:     "John",
			Email:    "john@example.com",
			Password: "password",
			Age:      25,
			Gender:   "Male",
			Session:  false,
		}
		user2 = model.User{
			Name:     "Jane",
			Email:    "jane@example.com",
			Password: "password",
			Age:      30,
			Gender:   "Female",
			Session:  false,
		}

		app.Reset()
	})

	Describe("LoginUser", func() {
		Context("when logging in with a valid email and password", func() {
			It("should set the user's session to true and return the user", func() {
				err := app.RegisterUser(user1)
				Expect(err).To(BeNil())

				u, err := app.LoginUser(user1.Email, user1.Password)
				Expect(err).To(BeNil())
				Expect(u.Session).To(BeTrue())

				app.Reset()
			})
		})

		Context("when logging in with an invalid email or password", func() {
			It("should return an error", func() {
				err := app.RegisterUser(user1)
				Expect(err).To(BeNil())

				_, err = app.LoginUser(user1.Email, "wrongpassword")
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("invalid email or password"))

				_, err = app.LoginUser("invalidemail@example.com", user1.Password)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("invalid email or password"))

				app.Reset()
			})
		})

	})

	Describe("GetLessonsByDifficulty", func() {
		Context("when the user is not logged in", func() {
			It("should return an error", func() {
				_, err := app.GetLessonsByDifficulty("john@example.com", 1)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("you must login first"))

				app.Reset()
			})
		})

		Context("when the user is logged in", func() {
			BeforeEach(func() {
				app.RegisterUser(user1)
				app.RegisterUser(user2)
				app.AddLesson(model.Lesson{
					Title:       "Introduction to Golang",
					Description: "Learn the basics of Go programming language",
					Category:    "Programming",
					Difficulty:  1,
				})
				app.AddLesson(model.Lesson{
					Title:       "Intermediate Golang",
					Description: "Learn advanced Go programming concepts",
					Category:    "Programming",
					Difficulty:  2,
				})
				app.LoginUser("john@example.com", "password")
			})

			It("should return lessons with the specified difficulty", func() {
				lessons, err := app.GetLessonsByDifficulty("john@example.com", 1)
				Expect(err).To(BeNil())
				Expect(lessons).To(HaveLen(1))
				Expect(lessons[0].Title).To(Equal("Introduction to Golang"))
				Expect(lessons[0].Difficulty).To(Equal(1))

				lessons, err = app.GetLessonsByDifficulty("john@example.com", 2)
				Expect(err).To(BeNil())
				Expect(lessons).To(HaveLen(1))
				Expect(lessons[0].Title).To(Equal("Intermediate Golang"))
				Expect(lessons[0].Difficulty).To(Equal(2))

				app.Reset()
			})

			It("should return an empty list if no lessons match the specified difficulty", func() {
				lessons, err := app.GetLessonsByDifficulty("john@example.com", 3)
				Expect(err).To(BeNil())
				Expect(lessons).To(HaveLen(0))

				app.Reset()
			})
		})
	})

})
