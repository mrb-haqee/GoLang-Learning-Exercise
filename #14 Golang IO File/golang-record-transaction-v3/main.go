package main

import (
	"bufio"
	"strings"

	// "errors"
	"fmt"
	"os"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

func RecordTransactions(path string, transactions []Transaction) error {
	//Counting Transaction
	Date := []string{}
	DateNow := ""
	for _, date := range transactions {
		check := false
		if DateNow != date.Date {
			for _, date2 := range Date {
				if date.Date == date2 {
					check = true
					break
				}
			}
			if !check {
				Date = append(Date, date.Date)
				DateNow = date.Date
			}
		}
	}
	DataFinal := []string{}
	for i := 0; i < len(Date); i++ {
		SumIncome := 0
		SumExpense := 0
		for j := 0; j < len(transactions); j++ {
			if transactions[j].Date == Date[i] {
				if transactions[j].Type == "income" {
					SumIncome += transactions[j].Amount
				} else if transactions[j].Type == "expense" {
					SumExpense += transactions[j].Amount
				}
			}
		}
		Total := SumIncome - SumExpense
		if Total > 0 {
			DataFinal = append(DataFinal, fmt.Sprintf("%s;income;%d", Date[i], Total))
		} else {
			Total = SumExpense - SumIncome
			DataFinal = append(DataFinal, fmt.Sprintf("%s;expense;%d", Date[i], Total))
		}

		//open file
		file, err := os.OpenFile(path, os.O_RDWR, 0644)
		isError(err)
		defer file.Close()

		//write file
		w := bufio.NewWriter(file)
		if len(Date) == 1 {
			for _, data := range DataFinal {
				_, err = w.WriteString(data)
				isError(err)
				w.Flush()
			}
		} else {
			for i, data := range DataFinal {
				if i == len(DataFinal)-1 {
					_, err = w.WriteString(data)
					isError(err)
					w.Flush()
					break
				}
				_, err = w.WriteString(fmt.Sprintf("%s\n", strings.TrimSpace(data)))
				isError(err)
				w.Flush()
			}
		}
	}
	return nil // TODO: replace this errors.New("not implemented")
}
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func main() {
	// bisa digunakan untuk pengujian test case
	var transactions = []Transaction{
		{"01/01/2021", "expense", 100000},
		{"01/01/2021", "expense", 1000},
		{"02/01/2021", "expense", 3424},
		{"02/01/2021", "expense", 42000},
		{"03/01/2021", "expense", 22321},
		{"04/01/2021", "expense", 223200},
		{"02/01/2021", "expense", 2300},
		{"05/01/2021", "expense", 2213},
		{"06/01/2021", "expense", 4545},
		{"07/01/2021", "expense", 55500},
		{"08/01/2021", "expense", 200000},
		{"10/01/2021", "expense", 20000},
		{"11/01/2021", "expense", 10000},
		{"12/01/2021", "expense", 55500},
		{"13/01/2021", "expense", 55500},
		{"02/01/2021", "expense", 55500},
		{"02/01/2021", "expense", 10000},
		{"14/01/2021", "expense", 20000},
		{"11/01/2021", "expense", 20000},
		{"15/01/2021", "expense", 10000},
		{"16/01/2021", "expense", 20000},
		{"02/01/2021", "expense", 55500},
		{"17/01/2021", "expense", 10000},
		{"06/01/2021", "expense", 20000},
		{"18/01/2021", "expense", 10000},
		{"03/01/2021", "expense", 20000},
		{"04/01/2021", "expense", 10000},
		{"19/01/2021", "expense", 55500},
		{"20/01/2021", "expense", 55500},
		{"21/01/2021", "expense", 10000},
		{"22/01/2021", "expense", 10000},
		{"23/01/2021", "expense", 10000},
		{"24/01/2021", "expense", 10000},
	}

	err := RecordTransactions("transactions.txt", transactions)
	isError(err)

	fmt.Println("Success")
}
