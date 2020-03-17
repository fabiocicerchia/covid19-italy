package userchoice

import (
	"testing"
	userchoice "../../../pkg/valueobject/userchoice"
)

func TestNewInstanceWithInvalidType(t *testing.T) {
	_, actual := userchoice.New("invalid", "test")

	expected := "The type \"invalid\" is not a valid type"
	if (actual != nil && actual.Error() != expected) {
		t.Errorf("Invalid error handling: %s", actual.Error())
		return
	}

	t.Logf("Correct error handling")
}

func TestNewInstanceWithInvalidValue(t *testing.T) {
	_, actual := userchoice.New("province", "")

	expected := "The value cannot be empty"
	if (actual != nil && actual.Error() != expected) {
		t.Errorf("Invalid error handling: %s", actual.Error())
		return
	}

	t.Logf("Correct error handling")
}

func TestNewInstanceWithProvinceValue(t *testing.T) {
	obj, err := userchoice.New("province", "roma")

	if err != nil {
		t.Errorf("No error should be returned: %s", err.Error())
	}

	if obj.GetSelectedType() == "province" {
		t.Logf("The SelectedType matches")
	} else {
		t.Errorf("The SelectedType doesn't match")
	}
	if obj.GetSelectedValue() == "roma" {
		t.Logf("The SelectedValue matches")
	} else {
		t.Errorf("The SelectedValuedoesn't match")
	}
}

func TestNewInstanceWithRegionValue(t *testing.T) {
	obj, err := userchoice.New("region", "lazio")

	if err != nil {
		t.Errorf("No error should be returned: %s", err.Error())
	}

	if obj.GetSelectedType() == "region" {
		t.Logf("The SelectedType matches")
	} else {
		t.Errorf("The SelectedType doesn't match")
	}
	if obj.GetSelectedValue() == "lazio" {
		t.Logf("The SelectedValue matches")
	} else {
		t.Errorf("The SelectedValuedoesn't match")
	}
}
