package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	number2 := numbers
	total := 0
	for number2 > 0 {
		number2 = number2 / 10
		total++
	}
	numString := strconv.Itoa(numbers)
	cast := []rune(numString)

	//ascii 0=48 9=57
	i := 0
	j := i + 1
	n, m := 0, 0   // variable menampung nilai int
	a, b := "", "" //Variable menampung nilai string
	jumlah1 := 0
	jumlah2 := jumlah1

	for i < total-1 {
		n, _ = strconv.Atoi(string(cast[i]))
		m, _ = strconv.Atoi(string(cast[j]))
		jumlah1 = n + m
		if jumlah1 > jumlah2 {
			jumlah2 = jumlah1
			a = strconv.Itoa(n)
			b = strconv.Itoa(m)
		} else if j == total-1 {
			break
		} else {
			i++;j++
		}
	}
	c := a + b
	Hasil, _ := strconv.Atoi(c)
	return Hasil // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(89083278))
}
