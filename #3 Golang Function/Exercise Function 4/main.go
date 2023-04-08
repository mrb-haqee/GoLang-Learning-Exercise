package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	check:=false
	DataCheck:=make([]string, 0)
	for i:=0; i<len(data); i++{
		check=strings.Contains(data[i], input)
		if check == true{
			DataCheck = append(DataCheck, data[i])
		}
	}
	SimilarData:=strings.Join(DataCheck, ",")

	return SimilarData // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
}
