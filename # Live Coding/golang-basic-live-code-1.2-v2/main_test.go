package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("ConvertLength", func() {
		When("with invalid input", func() {
			It("should return an error for negative input", func() {
				Expect(main.ConvertLength(-5, "m", "cm")).To(Equal(float64(0)))
			})

			It("should return an error for zero input", func() {
				Expect(main.ConvertLength(0, "m", "cm")).To(Equal(float64(0)))
			})
		})

		When("converting length from meter to other unit", func() {
			It("should return correct length in centimeter", func() {
				Expect(main.ConvertLength(1, "m", "cm")).To(Equal(100.0))
			})
			It("should return correct length in feet", func() {
				Expect(main.ConvertLength(1, "m", "ft")).To(Equal(3.281))
			})
			It("should return correct length in inch", func() {
				Expect(main.ConvertLength(1, "m", "in")).To(Equal(39.37))
			})
		})
		When("converting length from centimeter to other unit", func() {
			It("should return correct length in meter", func() {
				Expect(main.ConvertLength(100, "cm", "m")).To(Equal(1.0))
			})
			It("should return correct length in feet", func() {
				Expect(main.ConvertLength(100, "cm", "ft")).To(Equal(3.2808398950131235))
			})
			It("should return correct length in inch", func() {
				Expect(main.ConvertLength(100, "cm", "in")).To(Equal(39.37007874015748))
			})
		})
		When("converting length from feet to other unit", func() {
			It("should return correct length in meter", func() {
				Expect(main.ConvertLength(1, "ft", "m")).To(Equal(0.30478512648582745))
			})
			It("should return correct length in centimeter", func() {
				Expect(main.ConvertLength(1, "ft", "cm")).To(Equal(30.48))
			})
			It("should return correct length in inch", func() {
				Expect(main.ConvertLength(1, "ft", "in")).To(Equal(12.0))
			})
		})
		When("converting length from inch to other unit", func() {
			It("should return correct length in meter", func() {
				Expect(main.ConvertLength(1, "in", "m")).To(Equal(0.025400050800101603))
			})
			It("should return correct length in centimeter", func() {
				Expect(main.ConvertLength(1, "in", "cm")).To(Equal(2.54))
			})
			It("should return correct length in feet", func() {
				Expect(main.ConvertLength(1, "in", "ft")).To(Equal(0.08333333333333333))
			})
		})
		When("converting length to the same unit", func() {
			It("should return the same length", func() {
				Expect(main.ConvertLength(10, "m", "m")).To(Equal(10.0))
			})
		})
		When("converting length to an unsupported unit", func() {
			It("should return the same length", func() {
				Expect(main.ConvertLength(10, "m", "unsupported")).To(Equal(10.0))
			})
		})
	})
})
