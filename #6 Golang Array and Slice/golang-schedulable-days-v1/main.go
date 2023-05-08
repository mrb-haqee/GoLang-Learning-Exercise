package main

func SchedulableDays(date1 []int, date2 []int) []int {
	Holiday:=[]int{}
	for i:=0; i<len(date1); i++{
		for j:=0; j<len(date2); j++{
			if date1[i]==date2[j]{
				Holiday=append(Holiday, date1[i])
			}
		}
	}
	return Holiday // TODO: replace this
}
