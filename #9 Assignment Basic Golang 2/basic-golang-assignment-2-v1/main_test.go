package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("Test Login function", func() {
		When("given valid ID and name", func() {
			It("should return successful login message", func() {
				Expect(main.Login("A1234", "Aditira")).To(Equal("Login berhasil: Aditira"))
			})
		})

		When("given invalid ID or name", func() {
			When("ID is undefined", func() {
				It("should return error message", func() {
					Expect(main.Login("", "Aditira")).To(Equal("ID or Name is undefined!"))
				})
			})

			When("name is undefined", func() {
				It("should return error message", func() {
					Expect(main.Login("A1234", "")).To(Equal("ID or Name is undefined!"))
				})
			})

			When("data is not found", func() {
				It("should return error message", func() {
					Expect(main.Login("A1234", "Nanami")).To(Equal("Login gagal: data mahasiswa tidak ditemukan"))
				})
			})
		})
	})

	Describe("Test Register function", func() {
		When("given valid ID, name, and major", func() {
			It("should return successful registration message", func() {
				Expect(main.Register("C1112", "Yoga", "TI")).To(Equal("Registrasi berhasil: Yoga (TI)"))
				Expect(main.Students).To(HaveLen(4))
				Expect(main.Students).To(ContainElement("C1112_Yoga_TI"))
			})
		})

		When("given invalid ID, name, or major", func() {
			When("ID is undefined", func() {
				It("should return error message", func() {
					Expect(main.Register("", "Yoga", "TI")).To(Equal("ID, Name or Major is undefined!"))
				})
			})

			When("name is undefined", func() {
				It("should return error message", func() {
					Expect(main.Register("C1112", "", "TI")).To(Equal("ID, Name or Major is undefined!"))
				})
			})

			When("major is undefined", func() {
				It("should return error message", func() {
					Expect(main.Register("C1112", "Yoga", "")).To(Equal("ID, Name or Major is undefined!"))
				})
			})

			When("ID is already used", func() {
				It("should return error message", func() {
					Expect(main.Register("A1234", "Yoga", "TI")).To(Equal("Registrasi gagal: id sudah digunakan"))
					Expect(main.Students).To(HaveLen(4))
				})
			})
		})
	})

	Describe("Test GetStudyProgram", func() {
		It("should return Teknik Informatika", func() {
			Expect(main.GetStudyProgram("TI")).To(Equal("Teknik Informatika"))
		})

		It("should return Teknik Komputer", func() {
			Expect(main.GetStudyProgram("TK")).To(Equal("Teknik Komputer"))
		})

		It("should return Sistem Informasi", func() {
			Expect(main.GetStudyProgram("SI")).To(Equal("Sistem Informasi"))
		})

		It("should return Manajemen Informasi", func() {
			Expect(main.GetStudyProgram("MI")).To(Equal("Manajemen Informasi"))
		})

		It("should return Kode program studi tidak ditemukan", func() {
			Expect(main.GetStudyProgram("TA")).To(Equal("Kode program studi tidak ditemukan"))
		})
	})

	Describe("Test ModifyStudent", func() {
		It("should modify student's study program successfully", func() {
			programStudi := "GA"
			nama := "Dito"
			errMsg := main.ModifyStudent(programStudi, nama, main.UpdateStudyProgram)
			Expect(errMsg).To(Equal("Program studi mahasiswa berhasil diubah."))
			Expect(main.Students[1]).To(Equal("B2131_Dito_GA"))
		})

		It("should return 'Mahasiswa tidak ditemukan.' error message when the student is not found", func() {
			programStudi := "Sistem Informasi"
			nama := "Noname"
			errMsg := main.ModifyStudent(programStudi, nama, main.UpdateStudyProgram)
			Expect(errMsg).To(Equal("Mahasiswa tidak ditemukan."))
		})
	})
})
