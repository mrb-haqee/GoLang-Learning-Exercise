package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	Total := VIP + regular + student
	PriceVIP, PriceRegular, PriceStudent := 30, 20, 10
	TotalPrice := PriceVIP*VIP + PriceRegular*regular + PriceStudent*student
	if TotalPrice >= 100 {
		if day%2 != 0 {
			if Total < 5 {
				return (float32(TotalPrice) - float32(TotalPrice)*0.15)
			} else if Total >= 5 {
				return (float32(TotalPrice) - float32(TotalPrice)*0.25)
			}
		} else if day%2 == 0 {
			if Total < 5 {
				return (float32(TotalPrice) - float32(TotalPrice)*0.1)
			} else if Total >= 5 {
				return (float32(TotalPrice) - float32(TotalPrice)*0.2)
			}
		}
	} else {
		return float32(TotalPrice) // TODO: replace this
	}
	return float32(TotalPrice)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(4, 4, 4, 22))
}
