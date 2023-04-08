package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {
	avg := (math + science + english + indonesia) / 4
	if avg == 100 {
		return "Sempurna"
	} else if avg >= 90 {
		return "Sangat Baik"
	} else if avg >= 80 {
		return "Baik"
	} else if avg >= 70 {
		return "Cukup"
	} else if avg >= 60 {
		return "Kurang"
	} else if avg < 60 {
		return "Sangat kurang"
	}
	return "" // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 50, 50, 60))
}
