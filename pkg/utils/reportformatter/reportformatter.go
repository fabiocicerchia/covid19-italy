package reportformatter

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Output(choice string, data map[string]uint32) {
	fmt.Printf("COVID-19 NEW CASES IN %s\n\n", strings.ToUpper(choice))

	if len(data) == 0 {
		fmt.Print("No cases found\n")
		return
	}

	keys := sortedKeys(data)

	max := calculateMax(data)
	maxPadding := strconv.Itoa(len(strconv.Itoa(int(max))))
	worst := calculateWorst(data, keys)

	var perc int16 = 0
	var previous uint32 = 0
	var sumCases uint32 = 0
	for _, dataora := range keys {
		casi := uint32(data[dataora])
		sumCases = casi
		data := string(dataora[0:10])
		fmt.Printf("%s: %"+maxPadding+"d", data, int16(casi-previous))
		if previous > 0 {
			perc = int16(100.0/float32(previous)*float32(casi)) - 100
			fmt.Printf(" -> %+4d%%", perc)
		}
		if perc == worst && casi == max {
			fmt.Printf(" WORST INCREASE & PEAK")
		} else if perc == worst {
			fmt.Printf(" WORST INCREASE")
		} else if casi == max {
			fmt.Printf(" WORST PEAK")
		}
		fmt.Print("\n")
		previous = casi
	}

	fmt.Printf("\nTOTAL CASES: %d\n", sumCases)
}

func calculateMax(data map[string]uint32) uint32 {
	var max uint32 = 0
	for _, v := range data {
		if v > max {
			max = v
		}
	}

	return max
}

func calculateWorst(data map[string]uint32, keys []string) int16 {
	var max int16 = 0
	var previous uint32 = 0

	for _, k := range keys {
		casi := data[k]
		perc := int16(100.0/float32(previous)*float32(casi)) - 100
		previous = casi

		if perc > max {
			max = perc
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
