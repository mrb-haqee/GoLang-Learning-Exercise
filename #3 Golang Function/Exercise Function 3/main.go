package main

import (
	"fmt"
)

func ConvRune(JoinName [][]int) [][]string {
	result := make([][]string, len(JoinName))
	for i := 0; i < len(JoinName); i++ {
		for j := 0; j < len(JoinName[i]); j++ {
			result[i] = append(result[i], string(rune(JoinName[i][j])))
		}
	}
	return result
}

func FindShortestName(names string) string {
	//Make Symbol Rune
	Symbol := " ,;"
	SymbolRune := []rune(Symbol)
	numSymbolRune := []int{}
	for i := 0; i < len(SymbolRune); i++ {
		numSymbolRune = append(numSymbolRune, int(SymbolRune[i]))
	}
	//Symbol done

	//Make Name Rune
	namesRune := []rune(names)
	IntnamesRune := []int{}
	for i := 0; i < len(namesRune); i++ {
		IntnamesRune = append(IntnamesRune, int(namesRune[i]))
	}
	Nama := []int{}
	JoinName := [][]int{}
	count := 0
	for i := 0; i < len(IntnamesRune); i++ {
		if IntnamesRune[i] == numSymbolRune[0] || IntnamesRune[i] == numSymbolRune[1] || IntnamesRune[i] == numSymbolRune[2] {
			if i-count == 1 {
				for j := 0; j < i; j++ {
					Nama = append(Nama, IntnamesRune[j])
				}
			} else if i > count {
				for j := i - count + 1; j < i; j++ {
					Nama = append(Nama, IntnamesRune[j])
				}
			}
			JoinName = append(JoinName, Nama)
			Nama = []int{}
			count = 0
		} else if i == len(IntnamesRune)-1 {
			for j := i - count + 1; j <= i; j++ {
				Nama = append(Nama, IntnamesRune[j])
			}
			JoinName = append(JoinName, Nama)
			Nama = []int{}
		}
		if i >= 1 {
			count++
		}
	}
	//Name Rune Done output: Int-Rune

	//Name selection from length
	NameString := make([][]int, 0)
	PanjangName := 100
	for i := 0; i < len(JoinName); i++ {
		if len(JoinName[i]) <= PanjangName {
			if PanjangName == len(JoinName[i]) {
				NameString = append(NameString, JoinName[i])
			} else {
				PanjangName = len(JoinName[i])
				NameString = make([][]int, 0)
				NameString = append(NameString, JoinName[i])
			}
		}
	}
	// fmt.Println(ConvRune(NameString))
	//done Name selection from length

	//Comparation
	Hasil := ""
	if len(NameString) == 1 {
		NameString2 := ConvRune(NameString)
		for _, str := range NameString2[0] {
			Hasil += str
		}
	} else {

		LenName := len(NameString[0])
		LenJoinName := len(NameString)
		HasilRune := make([][]int, 0)
		Value := 123
		for i := 0; i < LenName; i++ {
			n := 0
			for j := 0; j < LenJoinName; j++ {
				if string(rune(NameString[j][i])) == "a" {
					HasilRune = make([][]int, 0)
					HasilRune = append(HasilRune, NameString[j])
					n=1
					break
				} else if j == LenJoinName-1 {
					if Value > 0 {
						n = 1
						break
					}
				} else if NameString[j][i] < Value {
					HasilRune = make([][]int, 0)
					Value = NameString[j][i]
					HasilRune = append(HasilRune, NameString[j])
				} else if NameString[j][i] > Value {
					continue
				} else {
					Value = 0
					HasilRune = make([][]int, 0)

				}
			}
			Value = 123
			if n == 1 {
				break
			}
		}
		HasilRune2 := ConvRune(HasilRune)
		for _, str := range HasilRune2[0] {
			Hasil += str
		}
	}
	//done Comparation
	return Hasil // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	// fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Ari,Aru,Ara,Andi,Asik")) // "Tia"
	// fmt.Println(FindShortestName("A,B,C,D,E")) // "Tia"
}
