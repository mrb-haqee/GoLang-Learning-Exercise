package main

import (
	"a21hc3NpZ25tZW50/helper"
	"fmt"
	"strings"
)

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"
var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"

func Login(id string, name string) string {
	// Make List Student as Array
	ListStudent := make([]string, 0)
	var before, after string
	for i := 0; i < 3; i++ {
		if i == 0 {
			before, after, _ = strings.Cut(Students, ", ")
			ListStudent = append(ListStudent, before)
		} else if i == 3-1 {
			ListStudent = append(ListStudent, after)
		} else {
			before, after, _ = strings.Cut(after, ", ")
			ListStudent = append(ListStudent, before)
		}
	}
	fmt.Println("ini List Student: ",ListStudent)
	//Done List Student
	////-----------------------------------------------------------------------
	//Make List ID, Name, Major as Array
	IDStudent := make([]string, 0)
	NameStudent := make([]string, 0)
	MajorStudent := make([]string, 0)
	for i := 0; i < len(ListStudent); i++ {
		before, after, _ = strings.Cut(ListStudent[i], "_")
		IDStudent = append(IDStudent, before)
		for {
			before, after, _ = strings.Cut(after, "_")
			NameStudent = append(NameStudent, before)
			MajorStudent = append(MajorStudent, after)
			break
		}
	}
	//Done list ID
	////-----------------------------------------------------------------------
	//Check Status ID and Name
	IdStatus := false
	ExID := 0
	for i := 0; i < len(IDStudent); i++ {
		IdStatus = strings.Contains(IDStudent[i], id)
		if IdStatus {
			ExID = i
			break
		}
	}
	NameStatus := false
	for i := 0; i < len(NameStudent); i++ {
		NameStatus = strings.Contains(NameStudent[i], name)
		if NameStatus {
			break
		}
	}
	//Check Done
	////-----------------------------------------------------------------------
	//Final Output
	var Status string
	if id == "" || name == "" {
		Status = "ID or Name is undefined!"
	} else if len(id) < 5 || len(id) > 5 {
		Status = "ID must be 5 characters long!"
	} else if IdStatus {
		if NameStatus {
			Status = fmt.Sprintf("Login berhasil: %s (%s)", NameStudent[ExID], MajorStudent[ExID])
		} else {
			Status = "Login gagal: data mahasiswa tidak ditemukan"
		}
	} else {
		Status = "Login gagal: data mahasiswa tidak ditemukan"
	}
	//Done!
	return Status // TODO: replace this
}

func Register(id string, name string, major string) string {
	// Make List Student as Array
	ListStudent := make([]string, 0)
	var before, after string
	for i := 0; i < 3; i++ {
		if i == 0 {
			before, after, _ = strings.Cut(Students, ", ")
			ListStudent = append(ListStudent, before)
		} else if i == 3-1 {
			ListStudent = append(ListStudent, after)
		} else {
			before, after, _ = strings.Cut(after, ", ")
			ListStudent = append(ListStudent, before)
		}
	}
	//Done List Student
	////-----------------------------------------------------------------------
	//Make List ID
	IDStudent := make([]string, 0)
	for i := 0; i < len(ListStudent); i++ {
		before, _, _ = strings.Cut(ListStudent[i], "_")
		IDStudent = append(IDStudent, before)
	}
	//Done list ID
	////-----------------------------------------------------------------------
	//Check Status ID
	IdStatus := false
	for i := 0; i < len(IDStudent); i++ {
		IdStatus = strings.Contains(IDStudent[i], id)
		if IdStatus {
			break
		}
	}
	//Check Done
	////-----------------------------------------------------------------------
	//Final Output
	var Status string
	if id == "" || name == "" || major == "" {
		Status = "ID, Name or Major is undefined!"
	} else if len(id) < 5 || len(id) > 5 {
		Status = "ID must be 5 characters long!"
	} else if IdStatus {
		Status = "Registrasi gagal: id sudah digunakan"
	} else {
		Status = fmt.Sprintf("Registrasi berhasil: %s (%s)", name, major)
	}
	//Done!
	return Status // TODO: replace this
}

func GetStudyProgram(code string) string {
	Major := ""
	if code == "TI" {
		Major = "Teknik Informatika"
	} else if code == "TK" {
		Major = "Teknik Komputer"
	} else if code == "SI" {
		Major = "Sistem Informasi"
	} else if code == "MI" {
		Major = "Manajemen Informasi"
	} else {
		Major = "Code is undefined!"
	}
	return Major // TODO: replace this
}

// func main() {
// 	fmt.Println(Login("B2131", "Juno"))
// }

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		fmt.Println("Students: ", Students)
		fmt.Println("Student Study Programs: ", StudentStudyPrograms)

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Keluar")

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
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
