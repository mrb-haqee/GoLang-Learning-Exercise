package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	var sm *main.InMemoryStudentManager

	BeforeEach(func() {
		sm = main.NewInMemoryStudentManager()
	})

	Describe("GetStudents", func() {
		It("should return all students", func() {
			students := sm.GetStudents()
			Expect(students).To(HaveLen(3))
			Expect(students[0].ID).To(Equal("A12345"))
			Expect(students[0].Name).To(Equal("Aditira"))
			Expect(students[0].StudyProgram).To(Equal("TI"))
			Expect(students[1].ID).To(Equal("B21313"))
			Expect(students[1].Name).To(Equal("Dito"))
			Expect(students[1].StudyProgram).To(Equal("TK"))
			Expect(students[2].ID).To(Equal("A34555"))
			Expect(students[2].Name).To(Equal("Afis"))
			Expect(students[2].StudyProgram).To(Equal("MI"))
		})
	})

	Describe("Login", func() {
		When("the ID and name match a student record", func() {
			It("should return success message", func() {
				msg, err := sm.Login("A12345", "Aditira")
				Expect(msg).To(Equal("Login berhasil: Aditira"))
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("the ID is empty", func() {
			It("should return an error", func() {
				msg, err := sm.Login("", "Aditira")
				Expect(msg).To(Equal(""))
				Expect(err).To(HaveOccurred())
			})
		})

		When("the name is empty", func() {
			It("should return an error", func() {
				msg, err := sm.Login("A12345", "")
				Expect(msg).To(Equal(""))
				Expect(err).To(HaveOccurred())
			})
		})

		When("the ID and name do not match any student record", func() {
			It("should return an error", func() {
				msg, err := sm.Login("invalid_id", "invalid_name")
				Expect(msg).To(Equal(""))
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Register", func() {
		When("all required fields are provided and ID is unique", func() {
			It("should add the student to the list", func() {
				msg, err := sm.Register("C12345", "Citra", "SI")
				Expect(msg).To(Equal("Registrasi berhasil: Citra (SI)"))
				Expect(err).ToNot(HaveOccurred())

				students := sm.GetStudents()
				Expect(students).To(HaveLen(4))
				Expect(students[3].ID).To(Equal("C12345"))
				Expect(students[3].Name).To(Equal("Citra"))
				Expect(students[3].StudyProgram).To(Equal("SI"))
			})
		})

		When("ID is already used", func() {
			It("should return an error", func() {
				msg, err := sm.Register("A12345", "Aditira", "TI")
				Expect(msg).To(Equal(""))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Registrasi gagal: id sudah digunakan"))
			})
		})

		When("ID, Name or StudyProgram is empty", func() {
			It("should return an error", func() {
				msg, err := sm.Register("", "", "TK")
				Expect(msg).To(Equal(""))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("ID, Name or StudyProgram is undefined!"))
			})
		})

		When("StudyProgram is invalid", func() {
			It("should return an error", func() {
				msg, err := sm.Register("C12345", "Citra", "ABC")
				Expect(msg).To(Equal(""))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Study program ABC is not found"))
			})
		})
	})

	Describe("GetStudyProgram", func() {
		When("given an existing study program code", func() {
			It("should return the correct study program name", func() {
				name, err := sm.GetStudyProgram("TI")
				Expect(err).NotTo(HaveOccurred())
				Expect(name).To(Equal("Teknik Informatika"))
			})
		})

		When("given an undefined study program code", func() {
			It("should return an error", func() {
				name, err := sm.GetStudyProgram("")
				Expect(err).To(HaveOccurred())
				Expect(name).To(BeEmpty())
			})
		})

		When("given a non-existing study program code", func() {
			It("should return an error", func() {
				name, err := sm.GetStudyProgram("unknown")
				Expect(err).To(HaveOccurred())
				Expect(name).To(BeEmpty())
			})
		})
	})

	Describe("ModifyStudent", func() {
		When("given an existing student name", func() {
			It("should modify the student's study program", func() {
				modifier := sm.ChangeStudyProgram("SI")

				msg, err := sm.ModifyStudent("Afis", modifier)
				Expect(err).NotTo(HaveOccurred())
				Expect(msg).To(Equal("Program studi mahasiswa berhasil diubah."))

				name, err := sm.GetStudyProgram("SI")
				Expect(err).NotTo(HaveOccurred())
				Expect(name).To(Equal("Sistem Informasi"))
			})
		})

		When("given a non-existing student name", func() {
			It("should return an error", func() {
				modifier := sm.ChangeStudyProgram("SI")

				msg, err := sm.ModifyStudent("unknown", modifier)
				Expect(err).To(HaveOccurred())
				Expect(msg).To(BeEmpty())
			})
		})
	})

})
