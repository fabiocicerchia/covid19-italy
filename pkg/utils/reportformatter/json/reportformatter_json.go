package reportformatter_json

import (
	"fmt"
	"sort"

	"encoding/json"
)

type ItemData struct {
	NewCases     uint32 `json:"new_cases"`
	TotalCases   uint32 `json:"total_cases"`
	PercIncrease uint32 `json:"perc_increase"`
	IsWorstPeak  bool   `json:"is_worst_peak"`
}

type OutputData struct {
	Choice string              `json:"choice"`
	Data   map[string]ItemData `json:"data"`
}

func Output(choice string, data map[string]uint32) {
	keys := sortedKeys(data)

	worst := calculateWorst(data, keys)

	var perc int16 = 0
	var previous uint32 = 0

	itemData := make(map[string]ItemData)

	for _, dataora := range keys {
		casi := uint32(data[dataora])
		data := string(dataora[0:10])
		newCases := int16(casi - previous)
		perc = 0
		if previous > 0 {
			perc = int16(100.0/float32(previous)*float32(casi)) - 100
		}
		itemData[data] = ItemData{
			NewCases:     uint32(newCases),
			TotalCases:   uint32(casi),
			PercIncrease: uint32(perc),
			IsWorstPeak:  newCases == worst,
		}

		previous = casi
	}

	outputData := OutputData{Choice: choice, Data: itemData}

	jsonString, _ := json.Marshal(outputData)
	fmt.Print(string(jsonString))
}

func calculateWorst(data map[string]uint32, keys []string) int16 {
	var max int16 = 0
	var previous uint32 = 0

	for _, k := range keys {
		newCases := int16(data[k] - previous)
		previous = data[k]

		if newCases > max {
			max = newCases
		}
	}

	return max
}

func sortedKeys(data map[string]uint32) []string {
	keys := make([]string, 0, len(data))
	for k, _ := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}
