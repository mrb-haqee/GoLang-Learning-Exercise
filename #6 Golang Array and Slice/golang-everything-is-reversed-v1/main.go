package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ReverseData(arr [5]int) [5]int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	rev := ""
	strRevSplit := []string{}
	strRev := []string{}
	for i := 0; i < len(arr); i++ {
		rev = fmt.Sprint(arr[i])
		strRevSplit = strings.Split(rev, "")
		for n, j := 0, len(strRevSplit)-1; n < j; n, j = n+1, j-1 {
			strRevSplit[n], strRevSplit[j] = strRevSplit[j], strRevSplit[n]
		}
		strRev = append(strRev, strings.Join(strRevSplit, ""))
	}

	numRev := 0
	FinalRev := []int{}
	for i := 0; i < len(strRev); i++ {
		numRev, _ = strconv.Atoi(strRev[i])
		FinalRev = append(FinalRev, numRev)
	}
	FinalRev2 := [5]int{}
	copy(FinalRev2[:], FinalRev)
	return FinalRev2 // TODO: replace this
}
