package reportfactory

import (
  "strconv"
  "os"

  "github.com/gocarina/gocsv"

  userchoice "cli/pkg/valueobject/userchoice"
  csv "cli/pkg/valueobject/csv"
  reportcasesservice "cli/pkg/service"
)

func Create(options map[string]string) *reportcasesservice.ReportCasesService {
    // TODO: Move elsewhere
    currPath, _ := os.Getwd()
    csvFileName := currPath + options["file"]

    csvFile, _:= os.OpenFile(csvFileName, os.O_RDONLY, os.ModePerm)
    defer csvFile.Close()

    rows := []*csv.BaseRow{}

    gocsv.UnmarshalFile(csvFile, &rows)

    // TODO: handle all the errors
    choice, err := userchoice.New(options["type"], options["userChoice"])
    if (err != nil) {
        panic(err)
    }
    lastDays, _ := strconv.Atoi(options["lastDays"])
    service := reportcasesservice.New(
        rows,
        *choice,
        lastDays,
    )

    return service
}
