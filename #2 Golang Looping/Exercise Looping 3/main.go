package main

import "fmt"

func CountingLetter(text string) int {
	// unreadable letters = R, S, T, Z
	hasil:=0
	n:=len(text)
	for i:=0; i<n; i++{
		if string(text[i])=="R"||string(text[i])=="S"||string(text[i])=="T"||string(text[i])=="Z"||string(text[i])=="r"||string(text[i])=="s"||string(text[i])=="t"||string(text[i])=="z"{
			hasil+=1
		}
	}
	return hasil // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Remaja muda yang berbakat"))
}
