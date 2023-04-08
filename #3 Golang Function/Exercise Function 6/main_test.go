package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MoneyChange", func() {
	When("input money 100000 and list price 50000, 10000, 10000, 5000, 5000", func() {
		It("should return 'Rp. 20.000'", func() {
			Expect(main.MoneyChange(100000, 50000, 10000, 10000, 5000, 5000)).To(Equal("Rp. 20.000"))
		})
	})

	When("input money 150000 and list price 50000, 50000, 50000, 30000", func() {
		It("should return 'Uang tidak cukup'", func() {
			Expect(main.MoneyChange(150000, 50000, 50000, 50000, 30000)).To(Equal("Uang tidak cukup"))
		})
	})

	When("input money 10000 and list price 5000, 1000", func() {
		It("should return 'Uang tidak cukup'", func() {
			Expect(main.MoneyChange(10000, 5000, 1000)).To(Equal("Rp. 4.000"))
		})
	})

	When("input money 10000 and list price 5000, 5000", func() {
		It("should return 'Uang tidak cukup'", func() {
			Expect(main.MoneyChange(10000, 5000, 5000)).To(Equal("Rp. 0"))
		})
	})

	When("input money 10000000 and list price 5000, 1000", func() {
		It("should return 'Uang tidak cukup'", func() {
			Expect(main.MoneyChange(10000000, 5000, 1000)).To(Equal("Rp. 9.994.000"))
		})
	})
})
