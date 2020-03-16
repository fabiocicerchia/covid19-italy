package getoptshandler

import (
    "fmt"
    "strings"
)

type GetoptsHandler struct {
    options map[string]string
}

type error interface {
    Error() string
}

func (o *GetoptsHandler) HandleOptions(region string, province string, lastWeek bool, lastMonth bool) (error) {
    o.options = make(map[string]string) // TODO: REMOVE?

    o.options["lastDays"] = "0"
    if (lastMonth) { o.options["lastDays"] = "31"
    } else if (lastWeek) { o.options["lastDays"] = "7" }

    if (region != "") {
        o.handlePlace("region", region);
    } else if (province != "") {
        o.handlePlace("province", province);
    }

    return nil
}

func (o GetoptsHandler) handlePlace(selectedType string, selectedValue string) (error) {
    if (selectedType == "region") {
        o.showRegion(selectedValue)
    } else if (selectedType == "province") {
        o.showProvince(selectedValue)
    } else {
        return fmt.Errorf("You must select a region or a province")
    }

    return nil
}

func (o *GetoptsHandler) showRegion(name string) {
    o.options["file"]       = "/data/pcm-dpc/dati-regioni/dpc-covid19-ita-regioni.csv"
    o.options["type"]       = "regione"
    o.options["userChoice"] = strings.ToLower(name)
}

func (o *GetoptsHandler) showProvince(name string) {
    o.options["file"]       = "/data/pcm-dpc/dati-province/dpc-covid19-ita-province.csv"
    o.options["type"]       = "provincia"
    o.options["userChoice"] = strings.ToLower(name)
}

func (o *GetoptsHandler) GetOptions() map[string]string {
    return o.options
}
