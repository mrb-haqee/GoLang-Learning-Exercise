package main

import (
	"fmt"
	"strconv"
)

func DateFormat(day, month, year int) string {
	CountDay := 0
	dayString := strconv.Itoa(day)
	yearString := strconv.Itoa(year)
	for day > 0 {
		day = day / 10
		CountDay++
	}
	Bulan := ""
	switch month {
	case 1:
		Bulan = "January"
	case 2:
		Bulan = "February"
	case 3:
		Bulan = "March"
	case 4:
		Bulan = "April"
	case 5:
		Bulan = "May"
	case 6:
		Bulan = "June"
	case 7:
		Bulan = "July"
	case 8:
		Bulan = "August"
	case 9:
		Bulan = "September"
	case 10:
		Bulan = "October"
	case 11:
		Bulan = "November"
	case 12:
		Bulan = "December"
	}
	fmt.Println(Bulan)
	Hasil := ""
	if CountDay == 1 {
		Hasil = "0" + dayString + "-" + Bulan + "-" + yearString
	} else {
		Hasil = dayString + "-" + Bulan + "-" + yearString
	}
	return Hasil // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
}
