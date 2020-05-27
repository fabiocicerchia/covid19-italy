package userchoice

import (
	"testing"

	"github.com/stretchr/testify/assert"

	userchoice "covid19.fabiocicerchia.it/cli/pkg/valueobject/userchoice"
)

func TestNewInstanceWithInvalidType(t *testing.T) {
	_, actual := userchoice.New("invalid", "test")

	expected := "The type \"invalid\" is not a valid type"

	assert.NotNil(t, actual)
	assert.Equal(t, actual.Error(), expected)
}

func TestNewInstanceWithInvalidValue(t *testing.T) {
	_, actual := userchoice.New("province", "")

	expected := "The value cannot be empty"

	assert.NotNil(t, actual)
	assert.Equal(t, actual.Error(), expected)
}

func TestNewInstanceWithProvinceValue(t *testing.T) {
	obj, err := userchoice.New("province", "roma")

	assert.Nil(t, err)

	assert.Equal(t, obj.GetSelectedType(), "province")
	assert.Equal(t, obj.GetSelectedValue(), "roma")
}

func TestNewInstanceWithRegionValue(t *testing.T) {
	obj, err := userchoice.New("region", "lazio")

	assert.Nil(t, err)

	assert.Equal(t, obj.GetSelectedType(), "region")
	assert.Equal(t, obj.GetSelectedValue(), "lazio")
}
