package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Readfile(path string) ([]string, error) {
	// membuka file
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil, err
	}
	defer file.Close()


	//membaca file 
	scanner := bufio.NewScanner(file)
	data := []string{}
	for scanner.Scan() {
		data = append(data, scanner.Text())
		if err := scanner.Err(); err != nil {
			fmt.Println("File reading error", err)
		}
	}
	return data, nil // TODO: replace this
}

func sum(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func CalculateProfitLoss(data []string) string {
	CheckIncome := false
	IncomeArrValue := []int{}
	ExpenseArrValue := []int{}
	CheckExpense := false
	for _, info := range data {
		rev := 0
		CheckIncome = strings.Contains(info, "income")
		CheckExpense = strings.Contains(info, "expense")
		if CheckIncome {
			rev, _ = strconv.Atoi(strings.Split(info, ";")[2])
			IncomeArrValue = append(IncomeArrValue, rev)
			CheckIncome = false
		}
		if CheckExpense {
			rev, _ = strconv.Atoi(strings.Split(info, ";")[2])
			ExpenseArrValue = append(ExpenseArrValue, rev)
			CheckExpense = false
		}
	}
	FinalValue := sum(IncomeArrValue) - sum(ExpenseArrValue)
	sort.Strings(data)
	DataSplit := strings.Split(data[len(data)-1], ";")
	if FinalValue > 0 {
		return DataSplit[0] + ";profit;" + fmt.Sprint(FinalValue)
	} else {
		return DataSplit[0] + ";loss;" + fmt.Sprint(int(math.Abs(float64(FinalValue)))) // TODO: replace this
	}
}

func main() {
	x, _ := Readfile("transactions.txt")
	fmt.Println(CalculateProfitLoss(x))
}
