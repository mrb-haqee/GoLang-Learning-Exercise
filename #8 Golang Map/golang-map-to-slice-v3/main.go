package main


func MapToSlice(mapData map[string]string) [][]string {
	DataSlice := make([][]string, len(mapData))
	n := 0
	for key, data := range mapData {
		DataSlice[n] = append(DataSlice[n], key)
		DataSlice[n] = append(DataSlice[n], data)
		n++
	}
	return DataSlice // TODO: replace this
}

