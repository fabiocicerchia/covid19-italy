package reportfactory

import (
	"os"

	"github.com/gocarina/gocsv"

	reportcasesservice "covid19.fabiocicerchia.it/cli/pkg/service/reportcasesservice"
	csv "covid19.fabiocicerchia.it/cli/pkg/valueobject/csv"
	userchoice "covid19.fabiocicerchia.it/cli/pkg/valueobject/userchoice"
)

func getCsvFileName(file string) string {
	currPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	csvFileName := currPath + file

	return csvFileName
}

func getCsvData(file string) []*csv.BaseRow {
	csvFile, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	rows := []*csv.BaseRow{}

	gocsv.UnmarshalFile(csvFile, &rows)

	return rows
}

func Create(csvFile string, selectedType string, selectedValue string, lastDays int) *reportcasesservice.ReportCasesService {
	csvFileName := getCsvFileName(csvFile)

	rows := getCsvData(csvFileName)

	choice, err := userchoice.New(selectedType, selectedValue)
	if err != nil {
		panic(err) // TODO: Handle all the panic
	}
	service := reportcasesservice.New(
		rows,
		*choice,
		lastDays,
	)

	return service
}
