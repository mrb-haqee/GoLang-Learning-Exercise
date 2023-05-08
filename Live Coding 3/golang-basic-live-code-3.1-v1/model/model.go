package model

type User struct {
	Name     string
	Email    string
	Password string
	Age      int
	Gender   string
	Session  bool
}

type Users []User

type Lesson struct {
	Title       string
	Description string
	Category    string
	Difficulty  int
}

type Lessons []Lesson
