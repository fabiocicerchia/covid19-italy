package reportfactory

import (
	"testing"

	"github.com/gocarina/gocsv"
	"github.com/stretchr/testify/assert"

	reportfactory "covid19.fabiocicerchia.it/cli/pkg/factory/reportfactory"
)

func TestProvinceData(t *testing.T) {
	gocsv.TagSeparator = ","

	csvFile := "/../../../data/pcm-dpc/dati-province/dpc-covid19-ita-province.csv"
	selectedType := "province"
	selectedValue := "roma"
	lastDays := 7

	service := reportfactory.Create(csvFile, selectedType, selectedValue, lastDays)

	assert.Equal(t, service.GetUserChoice().GetSelectedValue(), "roma")
}

func TestRegionData(t *testing.T) {
	csvFile := "/../../../data/pcm-dpc/dati-regioni/dpc-covid19-ita-regioni.csv"
	selectedType := "region"
	selectedValue := "lazio"
	lastDays := 7

	service := reportfactory.Create(csvFile, selectedType, selectedValue, lastDays)

	assert.Equal(t, service.GetUserChoice().GetSelectedValue(), "lazio")
}
