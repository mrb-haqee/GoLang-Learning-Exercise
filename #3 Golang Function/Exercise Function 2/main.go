package main

import (
	"fmt"
)

func CountVowelConsonant(str string) (int, int, bool) {
	strRune := []rune(str)
	numRune := []int{}
	for n := 0; n < len(strRune); n++ { //untuk str - merubah type rune menjadi int
		numRune = append(numRune, int(strRune[n]))
	}
	Vokal := "aiueoAIUEO"
	Syimbol := " ,'."
	VokalRune := []rune(Vokal)
	numVokalRune := []int{}
	SyimbolRune := []rune(Syimbol)
	numSyimbolRune := []int{}
	for n := 0; n < len(VokalRune); n++ { //untuk Vokal - merubah type rune menjadi int
		numVokalRune = append(numVokalRune, int(VokalRune[n]))
	}
	for n := 0; n < len(SyimbolRune); n++ { //untuk Syimbol - merubah type rune menjadi int
		numSyimbolRune = append(numSyimbolRune, int(SyimbolRune[n]))
	}

	addVokal := []int{}
	addSymbol := []int{}
	result:=true
	for n := 0; n < len(numVokalRune); n++ {
		for m := 0; m < len(numRune); m++ {
			if numVokalRune[n] == numRune[m] {
				addVokal = append(addVokal, numRune[m])
				result=false
			}
		}
	}
	for n := 0; n < len(numSyimbolRune); n++ {
		for m := 0; m < len(numRune); m++ {
			if numSyimbolRune[n] == numRune[m] {
				addSymbol = append(addSymbol, numRune[m])
			}
		}
	}
	LenumRune:=len(numRune)
	LVokal:=len(addVokal)
	LenumSymbolRune:=len(addSymbol)
	LConsonant:=LenumRune-LenumSymbolRune-LVokal

	return LVokal, LConsonant, result // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("bbbbb ccccc"))
}
