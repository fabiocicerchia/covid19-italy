package reportcasesservice

import (
	"strconv"
	"strings"

	csv "cli/pkg/valueobject/csv"
	userchoice "cli/pkg/valueobject/userchoice"
)

type ReportCasesService struct {
	file       []*csv.BaseRow
	lastDays   int
	userChoice userchoice.UserChoice
	data       map[string]uint32
}

func New(csv []*csv.BaseRow, userChoice userchoice.UserChoice, lastDays int) (obj *ReportCasesService) {
	obj = new(ReportCasesService)

	obj.file = csv
	obj.userChoice = userChoice
	obj.lastDays = lastDays

	return obj
}

func (s ReportCasesService) reverse(data []*csv.BaseRow) {
	for i := 0; i < len(data)/2; i++ {
		j := len(data) - i - 1
		data[i], data[j] = data[j], data[i]
	}
}

func (s *ReportCasesService) ProcessData() {
	// the file is sorted by date ascending, we need it descending
	// TODO: move it somewhere else
	csvData := s.file
	s.reverse(csvData)

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
