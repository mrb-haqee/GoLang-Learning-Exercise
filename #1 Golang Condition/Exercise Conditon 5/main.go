package main

import "fmt"

func tinggi(height int) int {
	if height > 160 {
		return 60000
	} else if height > 150 && height <= 160{
		return 40000
	} else if height > 135 && height <= 150{
		return 25000
	} else if height > 120 && height <= 135{
		return 15000
	}
	return 0
}

func TicketPlayground(height, age int) int {
	umur := age

	if age < 5 {
		return -1
	} else if age > 12 {
		return 100000
	} else {
		switch umur {
		case 5, 6, 7:
			if height > 120{
				return tinggi(height)
			} else{
				return 15000
			}
		case 8, 9:
			if height > 135{
				return tinggi(height)
			} else{
				return 25000
			}
		case 10, 11:
			if height > 150{
				return tinggi(height)
			} else{
				return 40000
			}
		case 12:
			if height > 160{
				return tinggi(height)
			} else{
				return 60000
			}
		}

	}

	return 0
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(151, 12))
}
