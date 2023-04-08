package main

import "fmt"

func CountingNumber(n int) float64 {
	total := 0.0
	for i := 1.0; i <= float64(n); i += 0.5 {
		total += i
	}
	return total
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(10))
}
