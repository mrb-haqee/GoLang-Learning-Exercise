package main

import (
	pack "a21hc3NpZ25tZW50/internal"
	"fmt"
	"strconv"
	"strings"
)

func AdvanceCalculator(calculate string) float32 {
	ArrData := strings.Split(calculate, " ")
	num1, _ := strconv.Atoi(ArrData[0])
	Hasil := pack.Calculator{Base: float32(num1)}
	for i:=0; i<len(ArrData); i++{
		if ArrData[i] == "*"{
			num, _ := strconv.Atoi(ArrData[i+1])
			Hasil.Multiply(float32(num))
		}else if ArrData[i] == "/"{
			num, _ := strconv.Atoi(ArrData[i+1])
			Hasil.Divide(float32(num))
		}else if ArrData[i] == "+"{
			num, _ := strconv.Atoi(ArrData[i+1])
			Hasil.Add(float32(num))
		}else if ArrData[i] == "-"{
			num, _ := strconv.Atoi(ArrData[i+1])
			Hasil.Subtract(float32(num))
		}
	}
	return Hasil.Result() // TODO: replace this
}

func main() {
	res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")
	fmt.Println(res)
}
