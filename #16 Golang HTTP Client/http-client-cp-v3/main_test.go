package main_test

import (
	main "a21hc3NpZ25tZW50"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("", func() {
	Describe("HTTP Client GET", func() {
		When("hit API https://animechan.vercel.app/api/quotes/anime?title=naruto and decode to struct Animechan", func() {
			It("should return struct Animechan with contents of the hit API", func() {
				res, err := main.ClientGet()
				Expect(err).To(BeNil())
				Expect(res).To(HaveLen(10))
				Expect(res[0].Anime).To(Equal("Naruto"))
				Expect(res[0].Character).To(Equal("Pain"))
				Expect(res[0].Quote).To(Equal("Because of the existence of love - sacrifice is born. As well as hate. Then one comprehends... one knows PAIN."))

				Expect(res[4].Anime).To(Equal("Naruto Shippuuden"))
				Expect(res[4].Character).To(Equal("Obito Uchiha"))
				Expect(res[4].Quote).To(Equal("The moment people come to know love, they run the risk of carrying hate."))

				Expect(res[9].Anime).To(Equal("Naruto Shippuuden"))
				Expect(res[9].Character).To(Equal("Madara Uchiha"))
				Expect(res[9].Quote).To(Equal("People cannot show each other their true feelings. Fear, suspicion, and resentment never subside."))
			})
		})
	})

	Describe("HTTP Client POST", func() {
		When("hit API https://postman-echo.com/post and decode to struct Postman", func() {
			It("should return struct Postman with contents of the hit API", func() {
				res, err := main.ClientPost()
				Expect(err).To(BeNil())
				Expect(res.Url).To(Equal("https://postman-echo.com/post"))
				Expect(res.Data.Name).To(Equal("Dion"))
				Expect(res.Data.Email).To(Equal("dionbe2022@gmail.com"))
			})
		})
	})
})
