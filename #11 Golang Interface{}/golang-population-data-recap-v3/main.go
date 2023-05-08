package main

import (
	"fmt"
	"strconv"
	"strings"
)

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
func PopulationData(data []string) []map[string]any {
	DataGeneral := make([][]string, len(data))
	DataGrand := map[string]any{}
	DataFinal := []map[string]any{}
	var (
		age       int
		height    float64
		isMarried bool
	)
	if len(data) > 1 {
		for i, info := range data {
			DataGeneral[i] = strings.Split(info, ";")
		}
		for _, ValueDataGeneral := range DataGeneral {
			if ValueDataGeneral[3] == "" && ValueDataGeneral[4] == "" {
				ValueDataGeneral = RemoveIndex(ValueDataGeneral, 4)
				ValueDataGeneral = RemoveIndex(ValueDataGeneral, 3)
				key := []string{"name", "age", "address"}
				DataGrand=map[string]any{}
				for i := 0; i < len(ValueDataGeneral); i++ {
					if i == 1 {
						age, _ = strconv.Atoi(ValueDataGeneral[i])
						DataGrand[key[i]] = age
						continue
					}
					DataGrand[key[i]] = ValueDataGeneral[i]
				}
				DataFinal = append(DataFinal, DataGrand)
			} else if ValueDataGeneral[3] == "" {
				ValueDataGeneral = RemoveIndex(ValueDataGeneral, 3)
				key := []string{"name", "age", "address", "isMarried"}
				DataGrand=map[string]any{}
				for i := 0; i < len(ValueDataGeneral); i++ {
					if i == 1 {
						age, _ = strconv.Atoi(ValueDataGeneral[i])
						DataGrand[key[i]] = age
						continue
					} else if i == 3 {
						isMarried, _ = strconv.ParseBool(ValueDataGeneral[i])
						DataGrand[key[i]] = isMarried
						break
					}
					DataGrand[key[i]] = ValueDataGeneral[i]
				}
				DataFinal = append(DataFinal, DataGrand)
			} else if ValueDataGeneral[4] == "" {
				ValueDataGeneral = RemoveIndex(ValueDataGeneral, 4)
				key := []string{"name", "age", "address", "height"}
				DataGrand=map[string]any{}
				for i := 0; i < len(ValueDataGeneral); i++ {
					if i == 1 {
						age, _ = strconv.Atoi(ValueDataGeneral[i])
						DataGrand[key[i]] = age
						continue
					} else if i == 3 {
						height, _ = strconv.ParseFloat(ValueDataGeneral[i], 64)
						DataGrand[key[i]] = height
						break
					} 
					DataGrand[key[i]] = ValueDataGeneral[i]
				}
				DataFinal = append(DataFinal, DataGrand)
			} else {
				key := []string{"name", "age", "address", "height", "isMarried"}
				DataGrand=map[string]any{}
				for i := 0; i < len(ValueDataGeneral); i++ {
					if i == 1 {
						age, _ = strconv.Atoi(ValueDataGeneral[i])
						DataGrand[key[i]] = age
						continue
					} else if i == 3 {
						height, _ = strconv.ParseFloat(ValueDataGeneral[i], 64)
						DataGrand[key[i]] = height
						continue
					} else if i == 4 {
						isMarried, _ = strconv.ParseBool(ValueDataGeneral[i])
						DataGrand[key[i]] = isMarried
						break
					}
					DataGrand[key[i]] = ValueDataGeneral[i]
				}
				DataFinal = append(DataFinal, DataGrand)
			}
		}
		return DataFinal
	} else {
		return []map[string]any{}
	}
}
func main() {
	data := []string{
		"Budi;23;Jakarta;160.42;true",
		"Joko;30;Bandung;;true",
		"Susi;25;Bogor;;false",
		"Santi;27;Jakarta;160;",
		"Rudi;23;Jakarta;161.1;",
		"Rudi;23;Jakarta;166.5;false",
		"Rudi;23;Jakarta;;",
	}
	fmt.Println("datanya: ")
	for _,info:=range PopulationData(data){
		fmt.Println(info)
	}
}
