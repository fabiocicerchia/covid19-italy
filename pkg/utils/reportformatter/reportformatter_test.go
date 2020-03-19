package reportformatter

import (
	"testing"
	"bytes"
	"io"
	"log"
	"os"
	"sync"

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

	expected := `COVID-19 TOTAL CASES IN ROMA
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

	expected := `COVID-19 TOTAL CASES IN ROMA
2020-03-10: 1
2020-03-11: 2 -> +100%
2020-03-12: 3 ->  +50%
2020-03-13: 4 ->  +33%
2020-03-14: 4 ->   +0%
2020-03-15: 3 ->  -25%
2020-03-16: 2 ->  -34%
2020-03-17: 1 ->  -50%
2020-03-18: 0 -> -100%
`

	assert.Equal(t, actual, expected)
}