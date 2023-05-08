package main

func Sortheight(height []int) []int {
	for i:=0; i<len(height); i++{
		for j:=0; j<len(height)-1-i;j++{
			if height[j]>height[j+1]{
				height[j], height[j+1]=height[j+1], height[j]
			}
		}
	}
	return height // TODO: replace this
}
