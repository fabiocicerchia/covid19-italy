package reportcasesservice

import (
	"testing"

	"github.com/gocarina/gocsv"
	"github.com/stretchr/testify/assert"

	reportcasesservice "cli/pkg/service/reportcasesservice"
	csv "cli/pkg/valueobject/csv"
	userchoice "cli/pkg/valueobject/userchoice"
)

func TestProcessDataProvince(t *testing.T) {
	// TODO: I'd like to do something like this:
	//csv := []*csv.BaseRow{}
	//csv = append(csv, csv.BaseRow{"2020-03-18", "Roma", "1"})
	// ...but got this one instead:
	csvContents := `data,denominazione_provincia,totale_casi
2020-03-10,roma,1
2020-03-10,milano,1
2020-03-11,roma,1
2020-03-11,milano,1
2020-03-12,roma,1
2020-03-12,milano,1
2020-03-13,roma,1
2020-03-13,milano,1
2020-03-14,roma,1
2020-03-14,milano,1
2020-03-15,roma,1
2020-03-15,milano,1
2020-03-16,roma,1
2020-03-16,milano,1
2020-03-17,roma,1
2020-03-17,milano,1
2020-03-18,roma,1
2020-03-18,milano,1
`
	csv := []*csv.BaseRow{}
	gocsv.UnmarshalString(csvContents, &csv)

	userChoice, _ := userchoice.New("province", "roma")
	lastDays := 7
	s := reportcasesservice.New(csv, *userChoice, lastDays)
	s.ProcessData()

	assert.Equal(t, len(s.GetData()), 7)
}

func TestProcessDataRegion(t *testing.T) {
	csvContents := `data,denominazione_regione,totale_casi
2020-03-10,lazio,1
2020-03-10,lombardia,1
2020-03-11,lazio,1
2020-03-11,lombardia,1
2020-03-12,lazio,1
2020-03-12,lombardia,1
2020-03-13,lazio,1
2020-03-13,lombardia,1
2020-03-14,lazio,1
2020-03-14,lombardia,1
2020-03-15,lazio,1
2020-03-15,lombardia,1
2020-03-16,lazio,1
2020-03-16,lombardia,1
2020-03-17,lazio,1
2020-03-17,lombardia,1
2020-03-18,lazio,1
2020-03-18,lombardia,1
`
	csv := []*csv.BaseRow{}
	gocsv.UnmarshalString(csvContents, &csv)

	userChoice, _ := userchoice.New("region", "lazio")
	lastDays := 7
	s := reportcasesservice.New(csv, *userChoice, lastDays)
	s.ProcessData()

	assert.Equal(t, len(s.GetData()), 7)
}
