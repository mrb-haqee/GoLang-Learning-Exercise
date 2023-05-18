package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type LoanData struct {
	StartBalance int
	Data         []Loan
	Employees    []Employee
}

type Loan struct {
	Date        string
	EmployeeIDs []string
}

type Employee struct {
	ID       string
	Name     string
	Position string
}

// json structure
type LoanRecord struct {
	MonthDate    string `json:"month_date"`
	StartBalance int `json:"start_balance"`
	EndBalance   int `json:"end_balance"`
	Borrowers    []Borrower `json:"borrowers"`
}

type Borrower struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Total int `json:"total_loan"`
}

func FindEmployee(id string, employees []Employee) (Employee, bool) {
	for _, employee := range employees {
		if employee.ID == id {
			return employee, true
		}
	}
	return Employee{}, false
}

func GetEndBalanceAndBorrowers(data LoanData) (int, []Borrower) {
	if len(data.Data) != 0 {
		// recap the data in map
		var tempMap = map[string]int{}
		for _, v := range data.Data {
			for _, id := range v.EmployeeIDs {
				if data.StartBalance >= 50000 {
					tempMap[id] += 50000
					data.StartBalance -= 50000
				}
			}
		}
		// sort temp
		var keys []string
		for k := range tempMap {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		var borrowers []Borrower
		// set data to struct
		for _, k := range keys {
			if employee, ok := FindEmployee(k, data.Employees); ok {
				borrowers = append(borrowers, Borrower{
					ID:    employee.ID,
					Name:  employee.Name,
					Total: tempMap[k],
				})
			}
		}

		// set the start balance and end balance
		return data.StartBalance, borrowers
	} else {
		return 0, []Borrower{}
	}
}

func LoanReport(data LoanData) LoanRecord {
	_, DataGetBorrower:=GetEndBalanceAndBorrowers(data)
	var DataFinal LoanRecord
	//Make MonthDate
	DateSplit := strings.Split(data.Data[0].Date, "-")
	DateJoin := strings.Join(DateSplit[1:], " ")
	DataFinal.MonthDate = DateJoin

	//StartBallance
	DataFinal.StartBalance = data.StartBalance

	//EndBalance
	StartBalance:=data.StartBalance
	for _,info:=range DataGetBorrower{
		StartBalance -=info.Total	
	}
	DataFinal.EndBalance=StartBalance

	//Borrower
	DataFinal.Borrowers=DataGetBorrower
	return DataFinal // TODO: replace this
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func RecordJSON(record LoanRecord, path string) error {
	//open file Json
	JsonFile, err := os.Open(path)
	isError(err)
	defer JsonFile.Close()

	records:=record
	byteValue, err :=json.Marshal(records)
	isError(err)
	err=ioutil.WriteFile(path, byteValue, 0644)
	isError(err)
	
	return nil // TODO: replace this
}

// gunakan untuk debug
func main() {
	records := LoanReport(LoanData{
		StartBalance: 500000,
		Data: []Loan{
			{
				Date:        "01-January-2021",
				EmployeeIDs: []string{"EMP001", "EMP002"},
			},
			{
				Date:        "02-January-2021",
				EmployeeIDs: []string{"EMP001", "EMP003"},
			},
		},
		Employees: []Employee{
			{
				ID:       "EMP001",
				Name:     "Eddy Assidiqi",
				Position: "Data Engineer",
			},
			{
				ID:       "EMP002",
				Name:     "Imam Permana",
				Position: "Frontend Engineer",
			},
			{
				ID:       "EMP003",
				Name:     "Rizky Ramadhan",
				Position: "Data Engineer",
			},
		},
	})

	err := RecordJSON(records, "loan-records.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(records)
}
