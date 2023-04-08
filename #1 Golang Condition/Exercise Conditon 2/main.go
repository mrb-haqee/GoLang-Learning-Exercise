package main

import "fmt"

func BMICalculator(gender string, height int) float64 {
	if gender == "laki-laki"{
		return ((float64(height)-100)-(float64(height)-100)*10/100)
	} else if gender == "perempuan"{
return ((float64(height)-100)-(float64(height)-100)*15/100)
	}
	return 0.0 // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	// fmt.Println(BMICalculator("laki-laki", 155))
	fmt.Println(BMICalculator("perempuan", 155))
}
