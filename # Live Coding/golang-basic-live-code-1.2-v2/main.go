package main

import (
	"fmt"
)

func main() {
	for {
		var panjang float64
		var dari, ke string

		fmt.Println("=== Kalkulator Konversi Satuan Panjang ===")

		fmt.Print("Masukkan panjang: ")
		fmt.Scanln(&panjang)

		fmt.Print("Masukkan satuan dari (m/cm/ft/in): ")
		fmt.Scanln(&dari)

		fmt.Print("Masukkan satuan ke (m/cm/ft/in): ")
		fmt.Scanln(&ke)

		result := ConvertLength(panjang, dari, ke)
		fmt.Printf("%.2f %s = %.2f %s\n", panjang, dari, result, ke)

		var pilihan string
		fmt.Print("Apakah Anda ingin mengkonversi kembali? (y/n): ")
		fmt.Scanln(&pilihan)

		if pilihan == "n" {
			break
		}
	}
}

func ConvertLength(panjang float64, dari, ke string) float64 {
	Convert := 0.0
	if panjang <= 0 {
		Convert = 0
	} else if panjang > 0 {
		if dari == "m" {
			if ke == "in" {
				Convert = 39.37 * panjang
			} else if ke == "cm" {
				Convert = 100.0 * panjang
			} else if ke == "ft" {
				Convert = 3.281 * panjang
			} else {
				Convert = panjang
			}
		} else if dari == "cm" {
			if ke == "in" {
				Convert = 0.3937007874015748 * panjang
			} else if ke == "m" {
				Convert = 0.01 * panjang
			} else if ke == "ft" {
				Convert = 0.032808398950131235 * panjang
			}else {
				Convert = panjang
			}
		} else if dari == "ft" {
			if ke == "in" {
				Convert = 12.0 * panjang
			} else if ke == "cm" {
				Convert = 30.48 * panjang
			} else if ke == "m" {
				Convert = 0.30478512648582745 * panjang
			}else {
				Convert = panjang
			}
		} else if dari == "in" {
			if ke == "m" {
				Convert = 0.025400050800101603 * panjang
			} else if ke == "cm" {
				Convert = 2.54 * panjang
			} else if ke == "ft" {
				Convert = 0.08333333333333333 * panjang
			}else {
				Convert = panjang
			}

		}
	}
	return Convert // TODO: replace this
}
