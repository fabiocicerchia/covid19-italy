package reportcasesservice

import (
  "strings"
  "strconv"

  userchoice "cli/pkg/valueobject/userchoice"
  csv "cli/pkg/valueobject/csv"
)

type ReportCasesService struct {
    file []*csv.BaseRow
    lastDays int
    userChoice userchoice.UserChoice
    data map[string]int
}

func New(csv []*csv.BaseRow, userChoice userchoice.UserChoice, lastDays int) *ReportCasesService {
    obj := ReportCasesService{}

    obj.file       = csv
    obj.userChoice = userChoice
    obj.lastDays   = lastDays

    return &obj
}

func (s ReportCasesService) reverse(data []*csv.BaseRow) []*csv.BaseRow {
    for i := 0; i < len(data) / 2; i++ {
        j := len(data) - i - 1
        data[i], data[j] = data[j], data[i]
    }

    return data
}

func (s *ReportCasesService) ProcessData() {
    // the file is sorted by date ascending, we need it descending
    // TODO: move it somewhere else
    csvData := s.reverse(s.file)

    data := make(map[string]int)

    countDays := 0
    for _, record := range csvData {
        if (strings.ToLower(record.Item) == strings.ToLower(s.userChoice.GetSelectedValue())) {
            cases, _ := strconv.Atoi(record.Cases)
            data[record.Data] = cases

            countDays++
            if (countDays == s.lastDays) {
                break
            }
        }
    }

    s.data = data;
}

func (s *ReportCasesService) GetData() map[string]int {
    return s.data;
}

func (s *ReportCasesService) GetUserChoice() userchoice.UserChoice {
    return s.userChoice;
}
