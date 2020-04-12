package reportformatter_json

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	reportformatter_json "cli/pkg/utils/reportformatter/json"
)

// Ref: https://gist.github.com/hauxe/e935a7f9012bf2649710cf75af323dbf
func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout, stderr := os.Stdout, os.Stderr
	defer func() {
		os.Stdout, os.Stderr = stdout, stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout, os.Stderr = writer, writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}

func TestOutputWithNoData(t *testing.T) {
	data := make(map[string]uint32)
	actual := captureOutput(func() {
		reportformatter_json.Output("roma", data)
	})

	expected := "{\"choice\":\"roma\",\"data\":{}}"

	assert.Equal(t, actual, expected)
}

func TestOutputWithSomeData(t *testing.T) {
	data := make(map[string]uint32)
	data["2020-03-10"] = 1
	data["2020-03-11"] = 2
	data["2020-03-12"] = 3
	data["2020-03-13"] = 4
	data["2020-03-14"] = 4
	data["2020-03-15"] = 3
	data["2020-03-16"] = 2
	data["2020-03-17"] = 1
	data["2020-03-18"] = 0

	actual := captureOutput(func() {
		reportformatter_json.Output("roma", data)
	})

	expected := "{\"choice\":\"roma\",\"data\":{\"2020-03-10\":{\"new_cases\":1,\"total_cases\":1,\"perc_increase\":0,\"is_worst_peak\":true},\"2020-03-11\":{\"new_cases\":1,\"total_cases\":2,\"perc_increase\":100,\"is_worst_peak\":true},\"2020-03-12\":{\"new_cases\":1,\"total_cases\":3,\"perc_increase\":50,\"is_worst_peak\":true},\"2020-03-13\":{\"new_cases\":1,\"total_cases\":4,\"perc_increase\":33,\"is_worst_peak\":true},\"2020-03-14\":{\"new_cases\":0,\"total_cases\":4,\"perc_increase\":0,\"is_worst_peak\":false},\"2020-03-15\":{\"new_cases\":4294967295,\"total_cases\":3,\"perc_increase\":4294967271,\"is_worst_peak\":false},\"2020-03-16\":{\"new_cases\":4294967295,\"total_cases\":2,\"perc_increase\":4294967262,\"is_worst_peak\":false},\"2020-03-17\":{\"new_cases\":4294967295,\"total_cases\":1,\"perc_increase\":4294967246,\"is_worst_peak\":false},\"2020-03-18\":{\"new_cases\":4294967295,\"total_cases\":0,\"perc_increase\":4294967196,\"is_worst_peak\":false}}}"

	assert.Equal(t, actual, expected)
}
