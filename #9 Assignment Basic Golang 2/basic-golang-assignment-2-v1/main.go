package main

import (
	"fmt"
	"strings"
	"a21hc3NpZ25tZW50/helper"
)

var Students = []string{
	"A1234_Aditira_TI",
	"B2131_Dito_TK",
	"A3455_Afis_MI",
}

var StudentStudyPrograms = map[string]string{
	"TI": "Teknik Informatika",
	"TK": "Teknik Komputer",
	"SI": "Sistem Informasi",
	"MI": "Manajemen Informasi",
}

type studentModifier func(string, *string)

func Login(id string, name string) string {
	IdCheck := false
	NameCheck:=false
	IndeksData := 0
	for i, data := range Students {
		IdCheck = strings.Contains((data), id)
		NameCheck = strings.Contains((data), name)
		if IdCheck {
			IndeksData = i
			break
		}
	}
	if id == "" || name == "" {
		return "ID or Name is undefined!"
	} else if IdCheck && NameCheck{
		DataStudent := []string{}
		DataStudent = append(DataStudent, strings.Split(Students[IndeksData], "_")...)
		return fmt.Sprintf("Login berhasil: %s", DataStudent[1])
	} else {
		return "Login gagal: data mahasiswa tidak ditemukan"
	}
	// TODO: replace this
}

func Register(id string, name string, major string) string {
	IdCheck := false
	for _, data := range Students {
		IdCheck = strings.Contains((data), id)
		if IdCheck {
			break
		}
	}
	if id == "" || name == "" || major == "" {
		return "ID, Name or Major is undefined!"
	} else if IdCheck {
		return "Registrasi gagal: id sudah digunakan"
	} else {
		Students=append(Students, fmt.Sprintf("%s_%s_%s", id, name, major))
		return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, major)
	}
	// TODO: replace this
}

func GetStudyProgram(code string) string {
	MajorCheck := false
	for key, data := range StudentStudyPrograms {
		MajorCheck = strings.Contains((key), strings.ToUpper(code))
		if MajorCheck {
			return data
		}
	}
	return "Kode program studi tidak ditemukan" // TODO: replace this
}

func ModifyStudent(programStudi, nama string, fn studentModifier) string {
	StudentCheck := false
	for i, data := range Students {
		StudentCheck = strings.Contains(data, nama)
		if StudentCheck {
			DataStudent := strings.Split(Students[i], "_")
			revisi := &DataStudent[2]
			fn(programStudi, revisi)
			Students[i] = strings.Join(DataStudent, "_")
			return "Program studi mahasiswa berhasil diubah."
		}
	}
	return "Mahasiswa tidak ditemukan." // TODO: replace this
}

func UpdateStudyProgram(programStudi string, students *string) {
	// fmt.Println("Masuk sini")
	*students = programStudi
	// TODO: answer here
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		for i, student := range Students {
			fmt.Println(i+1, student)
		}

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Change student study program")
		fmt.Println("5. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			helper.ClearScreen()
			var id, name string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)

			fmt.Println(Login(id, name))

			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			var id, name, jurusan string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)
			fmt.Print("Masukkan jurusan: ")
			fmt.Scan(&jurusan)
			fmt.Println(Register(id, name, jurusan))

			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			var kode string
			fmt.Print("Masukkan kode: ")
			fmt.Scan(&kode)

			fmt.Println(GetStudyProgram(kode))
			helper.Delay(5)
		case "4":
			helper.ClearScreen()
			var nama, programStudi string
			fmt.Print("Masukkan nama mahasiswa: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan program studi baru: ")
			fmt.Scanln(&programStudi)

			fmt.Println(ModifyStudent(programStudi, nama, UpdateStudyProgram))
			helper.Delay(5)
		case "5":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
