package main

import (
	"fmt"
	"strings"
)

func SlurredTalk(words *string) {
	checkS := strings.Contains(*words, "S")
	checkR := strings.Contains(*words, "R")
	checkZ := strings.Contains(*words, "Z")
	checks := strings.Contains(*words, "s")
	checkr := strings.Contains(*words, "r")
	checkz := strings.Contains(*words, "z")
	if checkS {
		*words = strings.ReplaceAll(*words, "S", "L")
	}
	if checkR {
		*words = strings.ReplaceAll(*words, "R", "L")
	}
	if checkZ {
		*words = strings.ReplaceAll(*words, "Z", "L")
	}
	if checks {
		*words = strings.ReplaceAll(*words, "s", "l")
	}
	if checkr {
		*words = strings.ReplaceAll(*words, "r", "l")
	}
	if checkz {
		*words = strings.ReplaceAll(*words, "z", "l")
	}
	// TODO: answer here
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Saya Steven saya suka menggoreng telur dan suka hewan zebra"
	SlurredTalk(&words)
	fmt.Println(words)
}
