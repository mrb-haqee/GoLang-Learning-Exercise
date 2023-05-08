package main

import (
	// "encoding/binary"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Time struct {
	Hour   int
	Minute int
}

func PrintType(x any) string {
	return reflect.TypeOf(x).String()
}

func ChangeToStandartTime(time any) string {
	tipe := PrintType(time)
	if tipe == "string" {
		SplitTime := strings.Split(time.(string), ":")
		if len(SplitTime) != 2 {
			return "Invalid input"
		}
		hour, _ := strconv.Atoi(SplitTime[0])
		if SplitTime[1] == "" {
			return "Invalid input"
		} else {
			if hour >= 12 {
				if hour == 12 {
					return fmt.Sprintf("12:%s PM", SplitTime[1])
				} else {
					hour = hour - 12
					if hour < 10 {
						return fmt.Sprintf("0%s:%s PM", strconv.Itoa(hour), SplitTime[1])
					} else {
						return fmt.Sprintf("%s:%s PM", strconv.Itoa(hour), SplitTime[1])
					}
				}
			} else {
				return fmt.Sprintf("%s:%s AM", SplitTime[0], SplitTime[1])
			}
		}
	} else if tipe == "[]int" {
		times := time.([]int)
		if len(times) <= 1 {
			return "Invalid input"
		}
		hour := times[0]
		minute := times[1]
		if hour >= 12 {
			if hour == 12 {
				return fmt.Sprintf("12:0%s PM", strconv.Itoa(minute))
			}
			hour = hour - 12
			if hour < 10 && minute < 10 {
				return fmt.Sprintf("0%s:0%s PM", strconv.Itoa(hour), strconv.Itoa(minute))
			} else if hour < 10 {
				return fmt.Sprintf("0%s:%s PM", strconv.Itoa(hour), strconv.Itoa(minute))
			} else if minute < 10 {
				return fmt.Sprintf("%s:0%s PM", strconv.Itoa(hour), strconv.Itoa(minute))
			} else {
				return fmt.Sprintf("%s:%s PM", strconv.Itoa(hour), strconv.Itoa(minute))
			}
		} else {
			if hour < 10 && minute < 10 {
				return fmt.Sprintf("0%s:0%s AM", strconv.Itoa(hour), strconv.Itoa(minute))
			} else if hour < 10 {
				return fmt.Sprintf("0%s:%s AM", strconv.Itoa(hour), strconv.Itoa(minute))
			} else if minute < 10 {
				return fmt.Sprintf("%s:0%s AM", strconv.Itoa(hour), strconv.Itoa(minute))
			} else {
				return fmt.Sprintf("%s:%s AM", strconv.Itoa(hour), strconv.Itoa(minute))
			}
		}
	} else if tipe == "map[string]int" {
		times := time.(map[string]int)
		hour := times["hour"]
		if len(times) <= 1 || hour==0{
			return "Invalid input"
		}
		minute := times["minute"]
		if hour >= 12 {
			if hour == 12 {
				return fmt.Sprintf("%s:0%s PM", strconv.Itoa(hour), strconv.Itoa(minute))
			}
			hour = hour - 12
			if hour < 10 && minute < 10 {
				return fmt.Sprintf("0%s:0%s PM", strconv.Itoa(hour), strconv.Itoa(minute))
			} else if hour < 10 {
				return fmt.Sprintf("0%s:%s PM", strconv.Itoa(hour), strconv.Itoa(minute))
			} else if minute < 10 {
				return fmt.Sprintf("%s:0%s PM", strconv.Itoa(hour), strconv.Itoa(minute))
			} else {
				return fmt.Sprintf("%s:%s PM", strconv.Itoa(hour), strconv.Itoa(minute))
			}
		} else {
			if hour < 10 && minute < 10 {
				return fmt.Sprintf("0%s:0%s AM", strconv.Itoa(hour), strconv.Itoa(minute))
			} else if hour < 10 {
				return fmt.Sprintf("0%s:%s AM", strconv.Itoa(hour), strconv.Itoa(minute))
			} else if minute < 10 {
				return fmt.Sprintf("%s:0%s AM", strconv.Itoa(hour), strconv.Itoa(minute))
			} else {
				return fmt.Sprintf("%s:%s AM", strconv.Itoa(hour), strconv.Itoa(minute))
			}
		}
	} else if tipe == "main.Time" {
		times := time.(Time)
		hour := times.Hour
		minute := times.Minute
		if hour >= 12 {
			if hour==12{
				return fmt.Sprintf("%s:0%s PM", strconv.Itoa(hour), strconv.Itoa(minute))
			}
			hour = hour - 12
			if hour < 10 && minute < 10 {
				return fmt.Sprintf("0%s:0%s PM", strconv.Itoa(hour), strconv.Itoa(minute))
			} else if hour < 10 {
				return fmt.Sprintf("0%s:%s PM", strconv.Itoa(hour), strconv.Itoa(minute))
			} else if minute < 10 {
				return fmt.Sprintf("%s:0%s PM", strconv.Itoa(hour), strconv.Itoa(minute))
			} else {
				return fmt.Sprintf("%s:%s PM", strconv.Itoa(hour), strconv.Itoa(minute))
			}
		} else {
			if hour < 10 && minute < 10 {
				return fmt.Sprintf("0%s:0%s AM", strconv.Itoa(hour), strconv.Itoa(minute))
			} else if hour < 10 {
				return fmt.Sprintf("0%s:%s AM", strconv.Itoa(hour), strconv.Itoa(minute))
			} else if minute < 10 {
				return fmt.Sprintf("%s:0%s AM", strconv.Itoa(hour), strconv.Itoa(minute))
			} else {
				return fmt.Sprintf("%s:%s AM", strconv.Itoa(hour), strconv.Itoa(minute))
			}
		}
	}
	return "" // TODO: replace this
}

func main() {
	fmt.Println(ChangeToStandartTime("2300"))
	fmt.Println(ChangeToStandartTime([]int{12, 00}))
	fmt.Println(ChangeToStandartTime(map[string]int{"hour": 16}))
	fmt.Println(ChangeToStandartTime(Time{16, 0}))

}
