package getoptshandler

import (
	"testing"

	"github.com/stretchr/testify/assert"

	getoptshandler "cli/pkg/utils/getoptshandler"
)

func TestHandleOptionsRegion(t *testing.T) {
	o := new(getoptshandler.GetoptsHandler)
	err := o.HandleOptions("test", "", false, false)

	assert.Nil(t, err)

	opts := o.GetOptions()

	assert.Equal(t, opts["file"], "/data/pcm-dpc/dati-regioni/dpc-covid19-ita-regioni.csv")
	assert.Equal(t, opts["type"], "region")
	assert.Equal(t, opts["userChoice"], "test")
}

func TestHandleOptionsProvince(t *testing.T) {
	o := new(getoptshandler.GetoptsHandler)
	err := o.HandleOptions("", "test", false, false)

	assert.Nil(t, err)

	opts := o.GetOptions()

	assert.Equal(t, opts["file"], "/data/pcm-dpc/dati-province/dpc-covid19-ita-province.csv")
	assert.Equal(t, opts["type"], "province")
	assert.Equal(t, opts["userChoice"], "test")
}

func TestHandleOptionsLastWeek(t *testing.T) {
	o := new(getoptshandler.GetoptsHandler)
	err := o.HandleOptions("test", "", false, true)

	assert.Nil(t, err)
	assert.Equal(t, o.GetOptions()["lastDays"], "7")
}

func TestHandleOptionsLastMonth(t *testing.T) {
	o := new(getoptshandler.GetoptsHandler)
	err := o.HandleOptions("test", "", true, false)

	assert.Nil(t, err)
	assert.Equal(t, o.GetOptions()["lastDays"], "31")
}

func TestHandleOptionsLastWeekAndLastMonth(t *testing.T) {
	o := new(getoptshandler.GetoptsHandler)
	err := o.HandleOptions("test", "", true, true)

	assert.Nil(t, err)
	assert.Equal(t, o.GetOptions()["lastDays"], "31")
}

func TestHandleOptionsNoRegionNorProvince(t *testing.T) {
	o := new(getoptshandler.GetoptsHandler)
	actual := o.HandleOptions("", "", false, false)

	expected := "You must select a region or a province"

	assert.NotNil(t, actual)
	assert.Equal(t, actual.Error(), expected)
}
