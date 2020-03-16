package reportfactory

import (
	"os"
	"strconv"

	"github.com/gocarina/gocsv"

	reportcasesservice "cli/pkg/service"
	csv "cli/pkg/valueobject/csv"
	userchoice "cli/pkg/valueobject/userchoice"
)

func Create(options map[string]string) *reportcasesservice.ReportCasesService {
	// TODO: Move elsewhere
	currPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	csvFileName := currPath + options["file"]

	csvFile, err := os.OpenFile(csvFileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	rows := []*csv.BaseRow{}

	gocsv.UnmarshalFile(csvFile, &rows)

	// TODO: handle all the errors
	choice, err := userchoice.New(options["type"], options["userChoice"])
	if err != nil {
		panic(err)
	}
	lastDays, err := strconv.Atoi(options["lastDays"])
	if err != nil {
		panic(err)
	}
	service := reportcasesservice.New(
		rows,
		*choice,
		lastDays,
	)

	return service
}
