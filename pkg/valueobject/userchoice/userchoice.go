package userchoice

import (
	"fmt"
)

type UserChoice struct {
	SelectedType  string
	SelectedValue string
}

type error interface {
	Error() string
}

func New(selectedType string, selectedValue string) (*UserChoice, error) {
	obj := new(UserChoice)

	if selectedType != "province" && selectedType == "region" {
		return obj, fmt.Errorf("The type \"%s\" is not a valid type", selectedType)
	}
	if selectedValue == "" {
		return obj, fmt.Errorf("The value cannot be empty")
	}

	obj.SelectedType = selectedType
	obj.SelectedValue = selectedValue

	return obj, nil
}

func (u UserChoice) GetSelectedType() string {
	return u.SelectedType
}

func (u UserChoice) GetSelectedValue() string {
	return u.SelectedValue
}
