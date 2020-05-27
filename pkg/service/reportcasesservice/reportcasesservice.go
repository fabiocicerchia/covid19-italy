package reportcasesservice

import (
	"strconv"
	"strings"

	csv "covid19.fabiocicerchia.it/cli/pkg/valueobject/csv"
	userchoice "covid19.fabiocicerchia.it/cli/pkg/valueobject/userchoice"
)

type ReportCasesService struct {
	csv        []*csv.BaseRow
	lastDays   int
	userChoice userchoice.UserChoice
	data       map[string]uint32
}

func New(csv []*csv.BaseRow, userChoice userchoice.UserChoice, lastDays int) (obj *ReportCasesService) {
	obj = new(ReportCasesService)

	obj.csv = csv
	obj.userChoice = userChoice
	obj.lastDays = lastDays

	return obj
}

func (s *ReportCasesService) ProcessData() {
	csvData := s.csv

	data := make(map[string]uint32)

	countDays := 0
	for _, record := range csvData {
		if strings.ToLower(record.Item) == strings.ToLower(s.userChoice.GetSelectedValue()) {
			cases, _ := strconv.Atoi(record.Cases)
			data[record.Data] = uint32(cases)

			countDays++
			if countDays == s.lastDays {
				break
			}
		}
	}

	s.data = data
}

func (s ReportCasesService) GetData() map[string]uint32 {
	return s.data
}

func (s ReportCasesService) GetUserChoice() userchoice.UserChoice {
	return s.userChoice
}
