package main

import "fmt"

// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {
	n := len(str)
	hasil := ""

	for i := n - 1; i >= 0; i-- {
		if string(str[i]) == " " {
			i--
			hasil += " " + string(str[i])
			continue
		} else if i == n-1 {
			hasil += string(str[i])
		} else {
			hasil += "_" + string(str[i])
		}
	}
	return hasil
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseString("Hello World"))
}
