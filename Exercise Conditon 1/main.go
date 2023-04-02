package main

import "fmt"

func GraduateStudent(score int, absent int) string {
	if score >= 70 && absent < 5 {
		return "lulus"
	} else if score < 70 && absent >= 5 {
		return "tidak lulus"
	} else {
		return "tidak lulus"
	}

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GraduateStudent(0, 6))
}
