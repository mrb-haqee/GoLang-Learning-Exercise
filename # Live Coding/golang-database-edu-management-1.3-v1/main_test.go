package main_test

import (
	main "a21hc3NpZ25tZW50"

	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Education Management", func() {
	var studentRepo repo.StudentRepository
	var teacherRepo repo.TeacherRepository

	dbCredential := model.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "mrb28",
		DatabaseName: "mrb_rg",
		Port:         5432,
		Schema:       "public",
	}
	conn, err := main.Connect(&dbCredential)
	Expect(err).ShouldNot(HaveOccurred())

	studentRepo = repo.NewStudentRepo(conn)
	teacherRepo = repo.NewTeacherRepo(conn)

	BeforeEach(func() {
		err = main.SQLExecute(conn)
		if err != nil {
			panic("failed droping table:" + err.Error())
		}

		err = main.Reset(conn, "students")
		err = main.Reset(conn, "teachers")
		Expect(err).ShouldNot(HaveOccurred())
	})

	Describe("Repository", func() {

		Describe("Student repository", func() {
			When("add student data to students table database postgres", func() {
				It("should save student data to students table database postgres", func() {
					student := model.Student{
						Name:    "John",
						Address: "Jl. Raya",
						Class:   "A",
					}
					err := studentRepo.Store(&student)
					Expect(err).ShouldNot(HaveOccurred())

					result, err := studentRepo.FetchByID(1)
					Expect(err).ShouldNot(HaveOccurred())

					Expect(result.Name).To(Equal(student.Name))
					Expect(result.Address).To(Equal(student.Address))
					Expect(result.Class).To(Equal(student.Class))

					err = main.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("read all student data from students table database postgres", func() {
				It("should return a list of student data", func() {
					student1 := model.Student{
						Name:    "John",
						Address: "Jl. Raya",
						Class:   "A",
					}
					err := studentRepo.Store(&student1)
					Expect(err).ShouldNot(HaveOccurred())

					student2 := model.Student{
						Name:    "Doe",
						Address: "Jl. Melati",
						Class:   "B",
					}
					err = studentRepo.Store(&student2)
					Expect(err).ShouldNot(HaveOccurred())

					result, err := studentRepo.FetchAll()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(result).To(HaveLen(2))
					Expect(result[0].Name).To(Equal(student1.Name))
					Expect(result[1].Name).To(Equal(student2.Name))

					err = main.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("adding new student data to students table in the database", func() {
				It("should save the new student data to students table in the database", func() {
					student := model.Student{
						Name:    "John",
						Address: "123 Main St",
						Class:   "Programming",
					}
					err := studentRepo.Store(&student)
					Expect(err).ShouldNot(HaveOccurred())

					result, err := studentRepo.FetchByID(1)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(result.Name).To(Equal(student.Name))
					Expect(result.Address).To(Equal(student.Address))
					Expect(result.Class).To(Equal(student.Class))

					err = main.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("fetching all student data from students table in the database", func() {
				Context("when there are no students in the database", func() {
					It("should return an empty slice", func() {
						result, err := studentRepo.FetchAll()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(len(result)).To(Equal(0))
					})
				})

				Context("when there are students in the database", func() {
					It("should return a list of all student data", func() {
						students := []model.Student{
							{Name: "John", Address: "123 Main St", Class: "Programming"},
							{Name: "Jane", Address: "456 Park Ave", Class: "Design"},
							{Name: "James", Address: "789 Broadway", Class: "Database"},
						}

						for _, student := range students {
							err := studentRepo.Store(&student)
							Expect(err).ShouldNot(HaveOccurred())
						}

						result, err := studentRepo.FetchAll()
						Expect(err).ShouldNot(HaveOccurred())
						Expect(len(result)).To(Equal(len(students)))

						for i, student := range students {
							Expect(result[i].Name).To(Equal(student.Name))
							Expect(result[i].Address).To(Equal(student.Address))
							Expect(result[i].Class).To(Equal(student.Class))
						}

						err = main.Reset(conn, "students")
						Expect(err).ShouldNot(HaveOccurred())
					})
				})
			})

			When("fetching a single student data by id from students table in the database", func() {
				It("should return a single student data", func() {
					student := model.Student{Name: "John", Address: "123 Main St", Class: "Programming"}
					err := studentRepo.Store(&student)
					Expect(err).ShouldNot(HaveOccurred())

					result, err := studentRepo.FetchByID(1)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(result.Name).To(Equal(student.Name))
					Expect(result.Address).To(Equal(student.Address))
					Expect(result.Class).To(Equal(student.Class))

					err = main.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("updating student data in students table in the database", func() {
				It("should update the existing student data in students table in the database", func() {
					student := model.Student{Name: "John", Address: "123 Main St", Class: "Programming"}
					err := studentRepo.Store(&student)
					Expect(err).ShouldNot(HaveOccurred())

					newStudent := model.Student{Name: "Jane", Address: "456 Park Ave", Class: "Design"}
					err = studentRepo.Update(1, &newStudent)
					Expect(err).ShouldNot(HaveOccurred())

					result, err := studentRepo.FetchByID(1)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(result.Name).To(Equal(newStudent.Name))
					Expect(result.Address).To(Equal(newStudent.Address))
					Expect(result.Class).To(Equal(newStudent.Class))

					err = main.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
		})
	})

	Describe("Teacher repository", func() {
		When("add teacher data to teachers table database postgres", func() {
			It("should save teacher data to teachers table database postgres", func() {
				teacher := model.Teacher{
					Name:    "John",
					Address: "Jl. Raya",
					Subject: "TI",
				}
				err := teacherRepo.Store(&teacher)
				Expect(err).ShouldNot(HaveOccurred())

				result, err := teacherRepo.FetchByID(1)
				Expect(err).ShouldNot(HaveOccurred())

				Expect(result.Name).To(Equal(teacher.Name))
				Expect(result.Address).To(Equal(teacher.Address))
				Expect(result.Subject).To(Equal(teacher.Subject))

				err = main.Reset(conn, "teachers")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		When("adding new teacher data to teachers table in the database", func() {
			It("should save the new teacher data to teachers table in the database", func() {
				teacher := model.Teacher{
					Name:    "John",
					Address: "123 Main St",
					Subject: "TI",
				}
				err := teacherRepo.Store(&teacher)
				Expect(err).ShouldNot(HaveOccurred())

				result, err := teacherRepo.FetchByID(1)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result.Name).To(Equal(teacher.Name))
				Expect(result.Address).To(Equal(teacher.Address))
				Expect(result.Subject).To(Equal(teacher.Subject))

				err = main.Reset(conn, "teachers")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		When("fetching a single teacher data by id from teachers table in the database", func() {
			It("should return a single teacher data", func() {
				teacher := model.Teacher{Name: "John", Address: "123 Main St", Subject: "TI"}
				err := teacherRepo.Store(&teacher)
				Expect(err).ShouldNot(HaveOccurred())

				result, err := teacherRepo.FetchByID(1)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result.Name).To(Equal(teacher.Name))
				Expect(result.Address).To(Equal(teacher.Address))
				Expect(result.Subject).To(Equal(teacher.Subject))

				err = main.Reset(conn, "teachers")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		When("deleting teacher data in teachers table in the database", func() {
			It("should delete the existing teacher data in teachers table in the database", func() {
				teacher := model.Teacher{Name: "John", Address: "123 Main St", Subject: "Programming"}
				err := teacherRepo.Store(&teacher)
				Expect(err).ShouldNot(HaveOccurred())

				err = teacherRepo.Delete(1)
				Expect(err).ShouldNot(HaveOccurred())

				result, err := teacherRepo.FetchByID(1)
				Expect(err).Should(HaveOccurred())
				Expect(result).To(BeNil())

				err = main.Reset(conn, "teachers")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})
})
