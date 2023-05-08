package main

import (
	"fmt"
	"sort"
)

type School struct {
	Name    string
	Address string
	Grades  []int
}

func (s *School) AddGrade(grades ...int) {
	for _, grade := range grades {
		if grade < 0 {
			grade = 0
		}
		if grade > 100 {
			grade = 100
		}
		s.Grades = append(s.Grades, grade)
	}
}

func Analysis(s School) (float64, int, int) {
	if len(s.Grades) == 0 {
		return 0.0, 0, 0
	} else {
		s.AddGrade(s.Grades...)
		sort.Ints(s.Grades)
		min := s.Grades[0]
		max := s.Grades[len(s.Grades)-1]
		total := 0
		for _, v := range s.Grades {
			total += v
		}
		avg := float64(total) / float64(len(s.Grades))
		return avg, min, max
	}

}

// gunakan untuk melakukan debugging
func main() {
	avg, min, max := Analysis(School{
		Name:    "Imam Assidiqi School",
		Address: "Jl. Imam Assidiqi",
		Grades:  []int{},
	})
	s:=School{}
	fmt.Println(avg, min, max)
	s.AddGrade()

}
