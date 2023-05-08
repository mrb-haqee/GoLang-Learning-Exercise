package main

import "fmt"

type Product struct {
	Name  string
	Price int
	Tax   int
}

func MoneyChanges(amount int, products []Product) []int {
	Total := 0
	for _, p := range products {
		Total += p.Price + p.Tax
	}
	Change := amount - Total
	GetCoin := []int{}
	for Change > 0 {
		if Change >= 1000 {
			Change = Change - 1000
			GetCoin = append(GetCoin, 1000)
		} else if Change >= 500 {
			Change = Change - 500
			GetCoin = append(GetCoin, 500)
		} else if Change >= 200 {
			Change = Change - 200
			GetCoin = append(GetCoin, 200)
		} else if Change >= 100 {
			Change = Change - 100
			GetCoin = append(GetCoin, 100)
		} else if Change >= 50 {
			Change = Change - 50
			GetCoin = append(GetCoin, 50)
		} else if Change >= 20 {
			Change = Change - 20
			GetCoin = append(GetCoin, 20)
		} else if Change >= 10 {
			Change = Change - 10
			GetCoin = append(GetCoin, 10)
		} else if Change >= 5 {
			Change = Change - 5
			GetCoin = append(GetCoin, 5)
		} else if Change >= 1 {
			Change--
			GetCoin = append(GetCoin, 1)
		} else {
			break
		}
	}
	return GetCoin // TODO: replace this
}
func main() {
	data := []Product{{Name: "Baju", Price: 5000, Tax: 500}, {Name: "Celana", Price: 3000, Tax: 300}}
	fmt.Println(MoneyChanges(10000, data))

	fmt.Printf("data: %v\n", data)
}
