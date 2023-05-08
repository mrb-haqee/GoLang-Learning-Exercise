package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PhoneNumberChecker(number string, result *string) {
	rev := ""
	check := false
	if len(number) < 11 {
		*result = "invalid"
	} else if string(number[0]) == "0" {
		for i := 811; i <= 888; i++ {
			if i >= 811 && i <= 815 {
				rev = "0" + strconv.Itoa(i)
				check = strings.Contains(number, rev)
				if check {
					*result = "Telkomsel"
					break
				}
				rev = ""
			}
			if i >= 815 && i <= 819 {
				rev = "0" + strconv.Itoa(i)
				check = strings.Contains(number, rev)
				if check {
					*result = "Indosat"
					break
				}
				rev = ""
			}
			if i >= 821 && i <= 823 {
				rev = "0" + strconv.Itoa(i)
				check = strings.Contains(number, rev)
				if check {
					*result = "XL"
					break
				}
				rev = ""
			}
			if i >= 827 && i <= 829 {
				rev = "0" + strconv.Itoa(i)
				check = strings.Contains(number, rev)
				if check {
					*result = "Tri"
					break
				}
				rev = ""
			}
			if i >= 852 && i <= 853 {
				rev = "0" + strconv.Itoa(i)
				check = strings.Contains(number, rev)
				if check {
					*result = "AS"
					break
				}
				rev = ""
			}
			if i >= 881 && i <= 888 {
				rev = "0" + strconv.Itoa(i)
				check = strings.Contains(number, rev)
				if check {
					*result = "Smartfren"
					break
				}
				rev = ""
			}
		}
		if check == false {
			*result = "invalid"
		}
	} else if string(number[0]) == "6" {
		for i := 62811; i <= 62888; i++ {
			if i >= 62811 && i <= 62815 {
				rev = strconv.Itoa(i)
				check = strings.Contains(number, rev)
				if check {
					*result = "Telkomsel"
					break
				}
				rev = ""
			}
			if i >= 62815 && i <= 62819 {
				rev = strconv.Itoa(i)
				check = strings.Contains(number, rev)
				if check {
					*result = "Indosat"
					break
				}
				rev = ""
			}
			if i >= 62821 && i <= 62823 {
				rev = strconv.Itoa(i)
				check = strings.Contains(number, rev)
				if check {
					*result = "XL"
					break
				}
				rev = ""
			}
			if i >= 62827 && i <= 62829 {
				rev = strconv.Itoa(i)
				check = strings.Contains(number, rev)
				if check {
					*result = "Tri"
					break
				}
				rev = ""
			}
			if i >= 62852 && i <= 62853 {
				rev = strconv.Itoa(i)
				check = strings.Contains(number, rev)
				if check {
					*result = "AS"
					break
				}
				rev = ""
			}
			if i >= 62881 && i <= 62888 {
				rev = strconv.Itoa(i)
				check = strings.Contains(number, rev)
				if check {
					*result = "Smartfren"
					break
				}
				rev = ""
			}
		}
		if check == false {
			*result = "invalid"
		}
	} else {
		*result = "invalid"
	}

	// TODO: answer here
}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "6281234"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}
