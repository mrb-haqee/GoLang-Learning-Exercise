package main

import (
	"a21hc3NpZ25tZW50/helper"
	"a21hc3NpZ25tZW50/model"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

type StudentManager interface {
	Login(id string, name string) error
	Register(id string, name string, studyProgram string) error
	GetStudyProgram(code string) (string, error)
	ModifyStudent(name string, fn model.StudentModifier) error
}

type InMemoryStudentManager struct {
	students             []model.Student
	studentStudyPrograms map[string]string
}

func NewInMemoryStudentManager() *InMemoryStudentManager {
	return &InMemoryStudentManager{
		students: []model.Student{
			{
				ID:           "A12345",
				Name:         "Aditira",
				StudyProgram: "TI",
			},
			{
				ID:           "B21313",
				Name:         "Dito",
				StudyProgram: "TK",
			},
			{
				ID:           "A34555",
				Name:         "Afis",
				StudyProgram: "MI",
			},
		},
		studentStudyPrograms: map[string]string{
			"TI": "Teknik Informatika",
			"TK": "Teknik Komputer",
			"SI": "Sistem Informasi",
			"MI": "Manajemen Informasi",
		},
	}
}

func (sm *InMemoryStudentManager) GetStudents() []model.Student {
	return sm.students // TODO: replace this
}

var (
	CheckID    bool
	CheckName  bool
	CheckMajor bool
)

func (sm *InMemoryStudentManager) Login(id string, name string) (string, error) {
	if id == "" || name == "" {
		return "", &CustomError{"ID or Name is undefined!"}
	} else {
		for _, InfoArr := range sm.students {
			CheckID = strings.Contains(InfoArr.ID, id)
			CheckName = strings.Contains(InfoArr.Name, name)
			if CheckID && CheckName {
				name = InfoArr.Name
				break
			}
		}
		if CheckID && CheckName {
			return fmt.Sprintf("Login berhasil: %s", name), nil
		} else {
			return "", &CustomError{"Login gagal: data mahasiswa tidak ditemukan"} // TODO: replace this
		}
	}
}

func (sm *InMemoryStudentManager) Register(id string, name string, studyProgram string) (string, error) {
	if id == "" || name == "" || studyProgram == "" {
		return "", &CustomError{"ID, Name or StudyProgram is undefined!"}
	} else {
		for _, InfoArr := range sm.students {
			CheckID = strings.Contains(InfoArr.ID, id)
			CheckName = strings.Contains(InfoArr.Name, name)
			if CheckID {
				break
			}
		}
		for key := range sm.studentStudyPrograms {
			CheckMajor = strings.Contains(key, studyProgram)
			if CheckMajor {
				break
			}
		}
		ErrorMajor := fmt.Sprintf("Study program %s is not found", studyProgram)
		if !CheckMajor {
			return "", &CustomError{ErrorMajor}
		} else if CheckID {
			return "", &CustomError{"Registrasi gagal: id sudah digunakan"}
		} else {
			sm.students = append(sm.students, model.Student{ID: id, Name: name, StudyProgram: studyProgram})
			return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, studyProgram), nil // TODO: replace this
		}
	}
}

func (sm *InMemoryStudentManager) GetStudyProgram(code string) (string, error) {
	if code == "" {
		return "", &CustomError{"Code is undefined!"}
	} else {
		for key := range sm.studentStudyPrograms {
			CheckMajor = strings.Contains(key, code)
			if CheckMajor {
				return sm.studentStudyPrograms[code], nil
			}
		}
		return "", &CustomError{"Kode program studi tidak ditemukan"} // TODO: replace this
	}
}

func (sm *InMemoryStudentManager) ModifyStudent(name string, fn model.StudentModifier) (string, error) {
	if name == "" {
		return "", &CustomError{"Mahasiswa tidak ditemukan."}
	} else {
		for _, InfoArr := range sm.students {
			CheckName := strings.Contains(InfoArr.Name, name)
			if CheckName {
				err := fn(&InfoArr)
				if err != nil {
					break
				}
				return "Program studi mahasiswa berhasil diubah.", nil // TODO: replace this
			}
		}
		return "", &CustomError{"Kode program studi tidak ditemukan"}
	}
}

func (sm *InMemoryStudentManager) ChangeStudyProgram(programStudi string) model.StudentModifier {
	for key := range sm.studentStudyPrograms {
		CheckMajor = strings.Contains(key, programStudi)
		if CheckMajor {
			return func(s *model.Student) error {
				s.StudyProgram = programStudi
				return nil
			}
		}
	}
	return func(s *model.Student) error {
		return &CustomError{"Kode program studi tidak ditemukan"}
	}
}

func main() {
	manager := NewInMemoryStudentManager()
	fmt.Println(manager.ModifyStudent("Aditira", manager.ChangeStudyProgram("adawdw")))
	for {
		helper.ClearScreen()
		students := manager.GetStudents()
		for _, student := range students {
			fmt.Printf("ID: %s\n", student.ID)
			fmt.Printf("Name: %s\n", student.Name)
			fmt.Printf("Study Program: %s\n", student.StudyProgram)
			fmt.Println()
		}

		fmt.Println("Selamat datang di Student Portal!")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Study Program")
		fmt.Println("4. Modify Student")
		fmt.Println("5. Exit")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Pilih menu: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			helper.ClearScreen()
			fmt.Println("=== Login ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			msg, err := manager.Login(id, name)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			fmt.Println("=== Register ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Study Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.Register(id, name, code)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			fmt.Println("=== Get Study Program ===")
			fmt.Print("Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			if studyProgram, err := manager.GetStudyProgram(code); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Printf("Program Studi: %s\n", studyProgram)
			}
			helper.Delay(5)
		case "4":
			helper.ClearScreen()
			fmt.Println("=== Modify Student ===")
			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Program Studi Baru (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.ModifyStudent(name, manager.ChangeStudyProgram(code))
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			helper.Delay(5)
		case "5":
			helper.ClearScreen()
			fmt.Println("Goodbye!")
			return
		default:
			helper.ClearScreen()
			fmt.Println("Pilihan tidak valid!")
			helper.Delay(5)
		}

		fmt.Println()
	}
}
