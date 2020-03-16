package getoptshandler

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

type GetoptsHandler struct {
	options map[string]string
}

type error interface {
	Error() string
}

func (o *GetoptsHandler) HandleOptions(c cli.Context) error {
	region := c.String("region")
	province := c.String("province")

	o.options = make(map[string]string)

	o.options["lastDays"] = "0"
	if c.Bool("last-month") {
		o.options["lastDays"] = "31"
	} else if c.Bool("last-week") {
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
	o.options["type"] = "regione"
	o.options["userChoice"] = strings.ToLower(name)
}

func (o *GetoptsHandler) showProvince(name string) {
	o.options["file"] = "/data/pcm-dpc/dati-province/dpc-covid19-ita-province.csv"
	o.options["type"] = "provincia"
	o.options["userChoice"] = strings.ToLower(name)
}

func (o GetoptsHandler) GetOptions() map[string]string {
	return o.options
}
