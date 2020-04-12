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

func (o *GetoptsHandler) HandleOptions(region string, province string, lastMonth bool, lastWeek bool) error {
	o.options = make(map[string]string)

	o.options["lastDays"] = "0"
	if lastMonth == true {
		o.options["lastDays"] = "31"
	} else if lastWeek == true {
		o.options["lastDays"] = "7"
	}

	if region != "" {
		return o.handlePlace("region", region)
	} else if province != "" {
		return o.handlePlace("province", province)
	}

	return fmt.Errorf("You must select a region or a province")
}

func (o GetoptsHandler) handlePlace(selectedType string, selectedValue string) error {
	if selectedType == "region" {
		o.showRegion(selectedValue)
	} else if selectedType == "province" {
		o.showProvince(selectedValue)
	} else {
		return fmt.Errorf("You must select a region or a province")
	}

	return nil
}

func (o *GetoptsHandler) showRegion(name string) {
	o.options["file"] = "/data/pcm-dpc/dati-regioni/dpc-covid19-ita-regioni.csv"
	o.options["type"] = "region"
	o.options["userChoice"] = strings.ToLower(name)
}

func (o *GetoptsHandler) showProvince(name string) {
	o.options["file"] = "/data/pcm-dpc/dati-province/dpc-covid19-ita-province.csv"
	o.options["type"] = "province"
	o.options["userChoice"] = strings.ToLower(name)
}

func (o GetoptsHandler) GetOptions() map[string]string {
	return o.options
}
