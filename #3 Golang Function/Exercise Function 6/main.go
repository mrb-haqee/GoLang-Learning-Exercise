package main

import (
	"fmt"
	"strconv"
)

func ChangeToCurrency(change int) string {
	//Don't worry about this implementation
	//Just use this function :)

	var res = ""

	moneyStr := strconv.Itoa(change)

	for i := len(moneyStr) - 1; i >= 0; i-- {
		if i == len(moneyStr)-1 {
			res = string(moneyStr[i]) + res
		} else if (len(moneyStr)-i)%3 == 0 {
			res = "." + string(moneyStr[i]) + res
		} else {
			res = string(moneyStr[i]) + res
		}
	}

	if res[0] == '.' {
		res = res[1:]
	}

	return "Rp. " + res
}

func MoneyChange(money int, listPrice ...int) string {
	for i:=0; i<len(listPrice); i++{
		money=money-listPrice[i]
		if money<0{
			return "Uang tidak cukup"
		}
	}
	return ChangeToCurrency(money) // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(MoneyChange(100_000, 50000, 10000, 10000, 5000, 5000))
	// fmt.Println(ChangeToCurrency(100_000))
}
