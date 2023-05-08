package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("SelectForm", func() {
		When("given a valid shape", func() {
			It("should return the shape if it is valid", func() {
				Expect(main.SelectForm("persegi")).To(Equal("persegi"))
				Expect(main.SelectForm("persegi-panjang")).To(Equal("persegi-panjang"))
			})
		})

		When("given an invalid shape", func() {
			It("should return an error message if the shape is invalid", func() {
				Expect(main.SelectForm("segitiga")).To(Equal("Bentuk geometri tidak valid!"))
				Expect(main.SelectForm("lingkaran")).To(Equal("Bentuk geometri tidak valid!"))
				Expect(main.SelectForm("jajargenjang")).To(Equal("Bentuk geometri tidak valid!"))
				Expect(main.SelectForm("bujur-sangkar")).To(Equal("Bentuk geometri tidak valid!"))
			})
		})
	})

	Describe("CalculateSquare", func() {
		When("with valid input", func() {
			It("should calculate the area and perimeter of a square", func() {
				luas, keliling, err := main.CalculateSquare(5)
				Expect(luas).To(Equal(25.0))
				Expect(keliling).To(Equal(10.0))
				Expect(err).To(Equal(""))
			})
		})

		When("with invalid input", func() {
			It("should return an error for negative input", func() {
				luas, keliling, err := main.CalculateSquare(-5)
				Expect(luas).To(Equal(float64(0)))
				Expect(keliling).To(Equal(float64(0)))
				Expect(err).To(Equal("sisi harus lebih besar dari 0"))
			})

			It("should return an error for zero input", func() {
				luas, keliling, err := main.CalculateSquare(0)
				Expect(luas).To(Equal(float64(0)))
				Expect(keliling).To(Equal(float64(0)))
				Expect(err).To(Equal("sisi harus lebih besar dari 0"))
			})
		})
	})

	Describe("CalculateRectangle", func() {
		When("with valid input", func() {
			It("should calculate the area and perimeter of a rectangle", func() {
				luas, keliling, err := main.CalculateRectangle(4, 6)
				Expect(luas).To(Equal(24.0))
				Expect(keliling).To(Equal(20.0))
				Expect(err).To(Equal(""))
			})
		})

		When("with invalid input", func() {
			It("should return an error for negative input", func() {
				luas, keliling, err := main.CalculateRectangle(-4, 6)
				Expect(luas).To(Equal(float64(0)))
				Expect(keliling).To(Equal(float64(0)))
				Expect(err).To(Equal("panjang dan lebar harus lebih besar dari 0"))
			})

			It("should return an error for zero input", func() {
				luas, keliling, err := main.CalculateRectangle(4, 0)
				Expect(luas).To(Equal(float64(0)))
				Expect(keliling).To(Equal(float64(0)))
				Expect(err).To(Equal("panjang dan lebar harus lebih besar dari 0"))
			})
		})
	})
})
