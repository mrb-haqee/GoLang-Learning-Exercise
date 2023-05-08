package main

import (
	"fmt"
)

func ViewStudents(students map[string][]interface{}) {
	fmt.Println("Name\tAddress\tPhone\tScore")
	for name, info := range students {
		address, _ := info[0].(string)
		phone, _ := info[1].(string)
		score, _ := info[2].(int)
		fmt.Printf("%s\t%s\t%s\t%d\n", name, address, phone, score)
	}
}

func AddStudent(students *map[string][]interface{}) func(string, string, string, int) {
	return func(name string, address string, phone string, score int) {
		(*students)[name] = []interface{}{address, phone, score}
	}
}

func RemoveStudent(students *map[string][]interface{}) func(string) {
	return func(name string) {
		delete(*students, name)
	} // TODO: replace this
}

func FindStudent(students map[string][]interface{}, scoreThreshold int) map[string][]interface{} {
	for name, info := range students {
		if info[2].(int) == scoreThreshold {
			murid:=map[string][]interface{}{}
			murid[name]=students[name]
			return murid
		}
	}
	return map[string][]interface{}{} // TODO: replace this
}

func main() {
	students := make(map[string][]interface{})
	add := AddStudent(&students)
	remove := RemoveStudent(&students)

	for {
		var command string
		fmt.Print("Enter command (add, remove, find, view): ")
		fmt.Scan(&command)

		switch command {
		case "add":
			var name, address, phone string
			var score int
			fmt.Print("Enter name: ")
			fmt.Scan(&name)
			fmt.Print("Enter address: ")
			fmt.Scan(&address)
			fmt.Print("Enter phone: ")
			fmt.Scan(&phone)
			fmt.Print("Enter score: ")
			fmt.Scan(&score)

			add(name, address, phone, score)
		case "remove":
			var name string
			fmt.Print("Enter name: ")
			fmt.Scan(&name)

			remove(name)
		case "find":
			var score int
			fmt.Print("Enter score threshold: ")
			fmt.Scan(&score)
			result := FindStudent(students, score)
			fmt.Println("Search result:")
			ViewStudents(result)
		case "view":
			fmt.Println("Student data:")
			ViewStudents(students)
		default:
			fmt.Println("Invalid command. Available commands: add, remove, find, view")
		}
	}
}
