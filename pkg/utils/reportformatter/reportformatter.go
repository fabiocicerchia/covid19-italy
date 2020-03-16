package reportformatter

import (
    "fmt"
    "sort"
    "strings"
    "strconv"
)

func Output(choice string, data map[string]int) {
    fmt.Printf("COVID-19 TOTAL CASES IN %s\n", strings.ToUpper(choice));

    if (len(data) == 0) {
        fmt.Print("No cases found\n")
        return
    }

    max := 0
    keys := make([]string, 0, len(data))
    for k, v := range data {
	keys = append(keys, k)
        if (v > max) { max = v }
    }
    sort.Strings(keys)
    maxPadding := strconv.Itoa(len(strconv.Itoa(max)))

    perc := 0
    previous := 0
    for _, dataora := range keys {
        casi := data[dataora]
        data := string(dataora[0:10])
        fmt.Printf("%s: %" + maxPadding + "d", data, casi);
        if (previous > 0) {
            perc = (100 / previous * casi) - 100
            fmt.Printf(" -> %+4d%%", perc);
        }
        fmt.Print("\n")
        previous = casi
    }
}
