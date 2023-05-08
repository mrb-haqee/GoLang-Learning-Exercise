package main

import (
	"strconv"
	"strings"
)

func SchedulableDays(villager [][]int) []int {
	FinalSchedul := []int{}
	if len(villager) == 0 {
		return FinalSchedul
	} else if len(villager) == 1 {
		return villager[0]
	} else {
		StrDayArr := []string{}
		for _, i := range villager[0] {
			StrDayArr = append(StrDayArr, strconv.Itoa(i))
		}
		StrDay := strings.Join(StrDayArr, ", ")
		StrDayArr = []string{}
		status := false
		for i := 1; i < len(villager); i++ {
			for j := 0; j < len(villager[i]); j++ {
				status = strings.Contains(StrDay, strconv.Itoa(villager[i][j]))
				if status {
					StrDayArr = append(StrDayArr, strconv.Itoa(villager[i][j]))
					status = false
				}
			}
			if i == len(villager)-1 {
				break
			}
			StrDay = strings.Join(StrDayArr, ", ")
		}
		rev := 0
		for _, i := range StrDayArr {
			rev, _ = strconv.Atoi(i)
			FinalSchedul = append(FinalSchedul, rev)
		}
		return FinalSchedul // TODO: replace this
	}
}

