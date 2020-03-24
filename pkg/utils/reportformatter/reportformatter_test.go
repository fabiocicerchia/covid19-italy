package reportformatter

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	reportformatter "cli/pkg/utils/reportformatter"
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
		reportformatter.Output("roma", data)
	})

	expected := `COVID-19 NEW CASES IN ROMA

No cases found
`

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
		reportformatter.Output("roma", data)
	})

	expected := `COVID-19 NEW CASES IN ROMA

DATE		NEW_CASES	TOTAL_CASES	% INCREASE
2020-03-10	1		1 WORST PEAK
2020-03-11	1		2 WORST PEAK
2020-03-12	1		3 WORST PEAK
2020-03-13	1		4 WORST PEAK
2020-03-14	0		4		   0%
2020-03-15	-1		3
2020-03-16	-1		2
2020-03-17	-1		1
2020-03-18	-1		0

TOTAL CASES: 0
`

	assert.Equal(t, actual, expected)
}
