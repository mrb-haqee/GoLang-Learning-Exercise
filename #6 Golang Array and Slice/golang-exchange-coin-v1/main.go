package main

func ExchangeCoin(amount int) []int {
	GetKoin := []int{}
	for amount >= 1_000 {
		amount = amount - 1_000
		GetKoin = append(GetKoin, 1_000)
	}
	for amount >= 500 {
		amount = amount - 500
		GetKoin = append(GetKoin, 500)
	}
	for amount >= 200 {
		amount = amount - 200
		GetKoin = append(GetKoin, 200)
	}
	for amount >= 100 {
		amount = amount - 100
		GetKoin = append(GetKoin, 100)
	}
	for amount >= 50 {
		amount = amount - 50
		GetKoin = append(GetKoin, 50)
	}
	for amount >= 20 {
		amount = amount - 20
		GetKoin = append(GetKoin, 20)
	}
	for amount >= 10 {
		amount = amount - 10
		GetKoin = append(GetKoin, 10)
	}
	for amount >= 5 {
		amount = amount - 5
		GetKoin = append(GetKoin, 5)
	}
	for amount >= 1 {
		amount = amount - 1
		GetKoin = append(GetKoin, 1)
	}
	return GetKoin // TODO: replace this
}
