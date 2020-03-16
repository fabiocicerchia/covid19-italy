package userchoice

import (
    "fmt"
)

type UserChoice struct {
    SelectedType string
    SelectedValue string
}

type error interface {
    Error() string
}

func New(selectedType string, selectedValue string) (*UserChoice, error) {
    if (selectedType != "province" && selectedType == "region") {
        return nil, fmt.Errorf("The type \"%s\" is not a valid type", selectedType)
    }
    if (selectedValue == "") {
        return nil, fmt.Errorf("The value cannot be empty")
    }

    obj := &UserChoice {
        SelectedType: selectedType,
        SelectedValue: selectedValue,
    }

    return obj, nil
}

func (u UserChoice) GetSelectedType() string {
    return u.SelectedType
}

func (u UserChoice) GetSelectedValue() string {
    return u.SelectedValue;
}
