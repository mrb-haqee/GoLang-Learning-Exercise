package main

import (
	"errors"
	"fmt"
	"strconv"
	// "strings"

	"a21hc3NpZ25tZW50/helper"
	"a21hc3NpZ25tZW50/model"
)

type Learnly interface {
	RegisterUser(user model.User) error
	LoginUser(email, password string) (model.User, error)
	GetLessonsByDifficulty(email string, difficulty int) ([]model.Lesson, error)
}

type learnlyApp struct {
	users   []model.User
	lessons []model.Lesson
}

func NewLearnly(users []model.User, lessons []model.Lesson) learnlyApp {
	return learnlyApp{users, lessons}
}

func (l *learnlyApp) GetUser() model.Users {
	return l.users
}

func (l *learnlyApp) AddLesson(lesson model.Lesson) {
	l.lessons = append(l.lessons, lesson)
}

func (l *learnlyApp) GetLesson() model.Lessons {
	return l.lessons
}

func (l *learnlyApp) Reset() {
	l.users = []model.User{}
	l.lessons = []model.Lesson{}
}

func (l *learnlyApp) Validate(u model.User) error {

	if u.Name == "" {
		return errors.New("name cannot be empty")
	}
	if u.Email == "" {
		return errors.New("email cannot be empty")
	}
	if u.Password == "" {
		return errors.New("password cannot be empty")
	}
	if u.Age < 0 || u.Age > 120 {
		return errors.New("age should be between 0 and 120")
	}
	if u.Gender != "Male" && u.Gender != "Female" {
		return errors.New("gender should be either Male or Female")
	}
	return nil
}

func (l *learnlyApp) RegisterUser(user model.User) error {
	if err := l.Validate(user); err != nil {
		return err
	}
	for _, u := range l.users {
		if u.Email == user.Email {
			return errors.New("email already registered")
		}
	}
	l.users = append(l.users, user)
	return nil
}

func (l *learnlyApp) LoginUser(email, password string) (model.User, error) {
	
	for _, u := range l.GetUser() {
        if u.Email == email && u.Password == password {
            u.Session = true
            return u, nil
        }
    }
    return model.User{}, errors.New("invalid email or password")
}



func (l *learnlyApp) GetLessonsByDifficulty(email string, difficulty int) ([]model.Lesson, error) {
	var foundUser *model.User
    for _, u := range l.GetUser() {
        if u.Email == email {
            foundUser = &u
        }
    }

    if foundUser == nil {
        return []model.Lesson{}, errors.New("you must login first")
    }

    if foundUser.Session {
        return []model.Lesson{}, errors.New("you must login first")
    }

    lessons := []model.Lesson{}
    for _, l := range l.GetLesson() {
        if l.Difficulty == difficulty {
            lessons = append(lessons, l)
        }
    }

    return lessons, nil
}

func main() {
	app := NewLearnly([]model.User{}, []model.Lesson{})

	var choice int
	for {
		helper.ClearScreen()
		fmt.Println("User: ", app.GetUser())
		fmt.Println("Lesson: ", app.GetLesson())

		fmt.Println("Welcome to LearnlyApp!")
		fmt.Println("1. Register user")
		fmt.Println("2. Login user")
		fmt.Println("3. Add lesson")
		fmt.Println("4. Get lesson by dificulty")
		fmt.Println("5. Exit")
		fmt.Print("Choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		switch choice {
		case 1:
			helper.ClearScreen()
			helper.InputScan("1. Register\n")
			name := helper.InputScan("\t- Name: ")
			email := helper.InputScan("\t- Email: ")
			password := helper.InputScan("\t- Password: ")
			age := helper.InputScan("\t- Age: ")
			ageInt, _ := strconv.Atoi(age)
			gender := helper.InputScan("\t- Gender (Male/Female): ")
			user := model.User{
				Name:     name,
				Email:    email,
				Password: password,
				Age:      ageInt,
				Gender:   gender,
				Session:  false,
			}
			err := app.RegisterUser(user)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("User registered successfully")
			}
			helper.Delay(5)
		case 2:
			helper.ClearScreen()
			helper.InputScan("2. Login\n")
			email := helper.InputScan("\t- Email: ")
			password := helper.InputScan("\t- Password: ")
			user, err := app.LoginUser(email, password)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Logged in as %s\n", user.Name)
			}
			helper.Delay(5)
		case 3:
			helper.ClearScreen()
			helper.InputScan("3. Add lesson\n")
			title := helper.InputScan("\t- Title: ")
			description := helper.InputScan("\t- Description: ")
			category := helper.InputScan("\t- Category: ")
			difficulty := helper.InputScan("\t- Difficulty: ")
			difficultyInt, _ := strconv.Atoi(difficulty)

			lesson := model.Lesson{
				Title:       title,
				Description: description,
				Category:    category,
				Difficulty:  difficultyInt,
			}

			app.AddLesson(lesson)
			helper.Delay(5)
		case 4:
			helper.ClearScreen()
			helper.InputScan("4. Get lesson by dificulty\n")
			email := helper.InputScan("\t- Email: ")
			difficulty := helper.InputScan("\t- Difficulty: ")
			difficultyInt, _ := strconv.Atoi(difficulty)

			res, err := app.GetLessonsByDifficulty(email, difficultyInt)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Lesson: ", res)
			}
			helper.Delay(5)
		case 5:
			fmt.Println("Thank you for using LearnlyApp!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
