package getoptshandler

import (
	"testing"

	getoptshandler "../../../pkg/utils/getoptshandler"
)

func TestHandleOptionsRegion(t *testing.T) {
	o := new(getoptshandler.GetoptsHandler)
	err := o.HandleOptions("test", "", false, false)

	if err != nil {
		t.Errorf("No error should be returned: %s", err.Error())
	}

	opts := o.GetOptions()
	if opts["file"] != "/data/pcm-dpc/dati-regioni/dpc-covid19-ita-regioni.csv" {
		t.Errorf("The type should be 'region': %s", opts["file"])
	}

	if opts["type"] != "region" {
		t.Errorf("The type should be 'region': %s", opts["type"])
	}

	if opts["userChoice"] != "test" {
		t.Errorf("The value should be 'test': %s", opts["userChoice"])
	}
}

func TestHandleOptionsProvince(t *testing.T) {
	o := new(getoptshandler.GetoptsHandler)
	err := o.HandleOptions("", "test", false, false)

	if err != nil {
		t.Errorf("No error should be returned: %s", err.Error())
	}

	opts := o.GetOptions()
	if opts["file"] != "/data/pcm-dpc/dati-province/dpc-covid19-ita-province.csv" {
		t.Errorf("The type should be 'region': %s", opts["file"])
	}

	if opts["type"] != "province" {
		t.Errorf("The type should be 'province': %s", opts["type"])
	}

	if opts["userChoice"] != "test" {
		t.Errorf("The value should be 'test': %s", opts["userChoice"])
	}
}

func TestHandleOptionsLastWeek(t *testing.T) {
	o := new(getoptshandler.GetoptsHandler)
	err := o.HandleOptions("test", "", false, true)

	if err != nil {
		t.Errorf("No error should be returned: %s", err.Error())
	}

	opts := o.GetOptions()
	if opts["lastDays"] != "7" {
		t.Errorf("The lastDays should be 7: %s", opts["lastDays"])
	}
}

func TestHandleOptionsLastMonth(t *testing.T) {
	o := new(getoptshandler.GetoptsHandler)
	err := o.HandleOptions("test", "", true, false)

	if err != nil {
		t.Errorf("No error should be returned: %s", err.Error())
	}

	if o.GetOptions()["lastDays"] != "31" {
		t.Errorf("The lastDays should be 31: %s", o.GetOptions()["lastDays"])
	}
}

func TestHandleOptionsLastWeekAndLastMonth(t *testing.T) {
	o := new(getoptshandler.GetoptsHandler)
	err := o.HandleOptions("test", "", true, true)

	if err != nil {
		t.Errorf("No error should be returned: %s", err.Error())
	}

	if o.GetOptions()["lastDays"] != "31" {
		t.Errorf("The lastDays should be 31: %s", o.GetOptions()["lastDays"])
	}
}

func TestHandleOptionsNoRegionNorProvince(t *testing.T) {
	o := new(getoptshandler.GetoptsHandler)
	actual := o.HandleOptions("", "", false, false)

	expected := "You must select a region or a province"
	if actual.Error() != expected {
		t.Errorf("Invalid error handling: %s", actual.Error())
	}
}
