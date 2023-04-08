package main

import (
	"fmt"
)

func FindMin(nums ...int) int {
	for i:=0; i<len(nums); i++{
		for j:=0; j<len(nums)-i-1; j++{
			if nums[j]>nums[j+1]{
				nums[j], nums[j+1]=nums[j+1], nums[j]
			}
		}
	}
	return nums[0] // TODO: replace this
}

func FindMax(nums ...int) int {
	for i:=0; i<len(nums); i++{
		for j:=0; j<len(nums)-i-1; j++{
			if nums[j]>nums[j+1]{
				nums[j], nums[j+1]=nums[j+1], nums[j]
			}
		}
	}
	return nums[len(nums)-1] // TODO: replace this
}

func SumMinMax(nums ...int) int {
	return FindMax(nums...)+FindMin(nums...) // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(SumMinMax(2,5,4,1,6,7,8,4,5,6,7,9,5,3,2,4,5,6,7,8,10))
}
