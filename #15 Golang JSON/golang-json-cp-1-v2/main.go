package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Study struct {
	StudyName   string `json:"study_name,omitempty"`
	StudyCredit int    `json:"study_credit"`
	Grade       string `json:"grade"`
}
type Report struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Date     string  `json:"date"`
	Semester int     `json:"semester"`
	Studies  []Study `json:"studies"` // TODO: answer here
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

// gunakan fungsi ini untuk mengambil data dari file json
// kembalian berupa struct 'Report' dan error
func ReadJSON(filename string) (Report, error) {
	//open file Json
	JsonFile, err := os.Open(filename)
	isError(err)
	defer JsonFile.Close()

	//read file Json
	byteValue, _ := ioutil.ReadAll(JsonFile)
	var Data Report
	err = json.Unmarshal(byteValue, &Data)
	isError(err)
	return Data, nil // TODO: answer here
}

func GradePoint(report Report) float64 {
	if len(report.Studies) <= 0 {
		return 0.0
	} else {
		var SumGrade float64
		var SumCredit float64
		for _, info := range report.Studies {
			if info.Grade == "A" {
				SumGrade += 4 * float64(info.StudyCredit)
			} else if info.Grade == "AB" {
				SumGrade += 3.5 * float64(info.StudyCredit)
			} else if info.Grade == "B" {
				SumGrade += 3 * float64(info.StudyCredit)
			} else if info.Grade == "BC" {
				SumGrade += 2.5 * float64(info.StudyCredit)
			} else if info.Grade == "C" {
				SumGrade += 2 * float64(info.StudyCredit)
			} else if info.Grade == "CD" {
				SumGrade += 1.5 * float64(info.StudyCredit)
			} else if info.Grade == "D" {
				SumGrade += 1 * float64(info.StudyCredit)
			} else if info.Grade == "DE" {
				SumGrade += 0.5 * float64(info.StudyCredit)
			} else if info.Grade == "E" {
				SumGrade += 0 * float64(info.StudyCredit)
			}
			SumCredit += float64(info.StudyCredit)
		}
		return SumGrade / SumCredit // TODO: replace this
	}
}

func main() {
	// bisa digunakan untuk menguji test case
	report, err := ReadJSON("report.json")
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
}
