package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Login", func() {
		When("successfully with registered student data", func() {
			It("Should return a success message", func() {
				actualMessage := main.Login("A1234", "Aditira")
				Expect(actualMessage).To(Equal("Login berhasil: Aditira (TI)"))
				actualMessage = main.Login("B2131", "Dito")
				Expect(actualMessage).To(Equal("Login berhasil: Dito (TK)"))
				actualMessage = main.Login("A3455", "Afis")
				Expect(actualMessage).To(Equal("Login berhasil: Afis (MI)"))
			})
		})

		When("fails with unregistered student data", func() {
			It("Should return a failure message", func() {
				id := "B1234"
				name := "Bambang"
				expectedMessage := "Login gagal: data mahasiswa tidak ditemukan"
				actualMessage := main.Login(id, name)
				Expect(actualMessage).To(Equal(expectedMessage))
			})
		})

		When("fails with incomplete student data", func() {
			It("Should return a failure message", func() {
				id := "A1238"
				name := "Adit"
				expectedMessage := "Login gagal: data mahasiswa tidak ditemukan"
				actualMessage := main.Login(id, name)
				Expect(actualMessage).To(Equal(expectedMessage))
			})
		})

		When("fails with incorrect student data", func() {
			It("Should return a failure message", func() {
				id := "B2131"
				name := "Juno"
				expectedMessage := "Login gagal: data mahasiswa tidak ditemukan"
				actualMessage := main.Login(id, name)
				Expect(actualMessage).To(Equal(expectedMessage))
			})
		})

		When("with an empty ID", func() {
			It("should return error message", func() {
				result := main.Login("", "Dina")
				Expect(result).To(Equal("ID or Name is undefined!"))
			})
		})

		When("with an ID less or more than 5", func() {
			It("should return error message", func() {
				result := main.Login("A123", "Aditira")
				Expect(result).To(Equal("ID must be 5 characters long!"))
				result = main.Login("A12345", "Aditira")
				Expect(result).To(Equal("ID must be 5 characters long!"))
			})
		})

		When("with an empty Name", func() {
			It("should return error message", func() {
				result := main.Login("A1234", "")
				Expect(result).To(Equal("ID or Name is undefined!"))
			})
		})
	})

	Describe("Register", func() {
		When("registering a new student with a unique ID", func() {
			It("should return success message", func() {
				result := main.Register("C1234", "Cindy", "TI")
				Expect(result).To(Equal("Registrasi berhasil: Cindy (TI)"))
			})
		})

		When("registering a new student with an ID that already exists", func() {
			It("should return error message", func() {
				result := main.Register("A1234", "Ali", "TK")
				Expect(result).To(Equal("Registrasi gagal: id sudah digunakan"))
			})
		})

		When("registering a new student with an empty name", func() {
			It("should return error message", func() {
				result := main.Register("D1234", "", "SI")
				Expect(result).To(Equal("ID, Name or Major is undefined!"))
			})
		})

		When("registering a new student with an empty ID", func() {
			It("should return error message", func() {
				result := main.Register("", "Diana", "MI")
				Expect(result).To(Equal("ID, Name or Major is undefined!"))
			})
		})

		When("registering a new student with an ID less or more than 5", func() {
			It("should return error message", func() {
				result := main.Register("A123", "Aditira", "MI")
				Expect(result).To(Equal("ID must be 5 characters long!"))
				result = main.Register("A12345", "Aditira", "MI")
				Expect(result).To(Equal("ID must be 5 characters long!"))
			})
		})

		When("registering a new student with an empty jurusan", func() {
			It("should return error message", func() {
				result := main.Register("E1234", "Erika", "")
				Expect(result).To(Equal("ID, Name or Major is undefined!"))
			})
		})
	})

	Describe("GetStudyProgram", func() {
		When("given an empty code", func() {
			It("should return 'Code is undefined!'", func() {
				Expect(main.GetStudyProgram("")).To(Equal("Code is undefined!"))
			})
		})

		When("given a valid code", func() {
			It("should return the corresponding study program", func() {
				Expect(main.GetStudyProgram("TI")).To(Equal("Teknik Informatika"))
				Expect(main.GetStudyProgram("TK")).To(Equal("Teknik Komputer"))
				Expect(main.GetStudyProgram("SI")).To(Equal("Sistem Informasi"))
				Expect(main.GetStudyProgram("MI")).To(Equal("Manajemen Informasi"))
			})
		})
	})
})
