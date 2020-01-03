package glog

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

/*
 * TestLog test the log function (to a file output)
 * tested Log()
 * tested Mute()
 */
func TestLog(t *testing.T) {
	logFile := "./info.log"
	i := INFO
	mockTime := "2020-01-01 00:00:00 +0000 UTC"

	// Setup the log file and the logger
	// --------------------------------------------------------------------
	// Remove log file if exist
	os.Remove(logFile)

	// Mock time for assess output

	out, err := NewFileOutput("./info.log")
	if err == nil {
		i.To(out, "log file")
	}

	// Change format function to print mock time
	i.Format = func(l *Log, logMsg ...interface{}) string {
		fmtLog := []interface{}{mockTime, l.Flag, fmt.Sprint(logMsg...)}
		log := fmt.Sprintf("%v [%v] - %v\n", fmtLog...)
		return log
	}
	// --------------------------------------------------------------------

	// Log hello word
	// Output: 2020-01-01 00:00:00 +0000 UTC [INFO] - Hello world
	i.Log("Hello world")
	// Mute
	i.Mute()
	// Shouldn't be logged
	i.Log("Bye world")

	// Check the log file
	// --------------------------------------------------------------------
	f, err := os.Open(logFile)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	// Store the line of the log file
	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err != nil {
		t.Fatal(err)
	}

	// Mute not working, more or less than 1 log writted
	if len(lines) != 1 {
		t.Fatalf("number of line expected : 1 \nnumber of line found : %v", len(lines))
	}

	// Check the log format
	output := "2020-01-01 00:00:00 +0000 UTC [INFO] - Hello world"
	if lines[0] != output {
		t.Fatalf("Expected output : \n%v \ngot : %v", output, lines[0])
	}
	// --------------------------------------------------------------------
}

/*
 * TestLogf test the logf function (to a file output)
 * tested Logf()
 * tested Unmute()
 */
func TestLogf(t *testing.T) {
	logFile := "./info.log"
	i := INFO
	mockTime := "2020-01-01 00:00:00 +0000 UTC"

	// Setup the log file and the logger
	// --------------------------------------------------------------------
	// Remove log file if exist
	os.Remove(logFile)

	// Mock time for assess output

	out, err := NewFileOutput("./info.log")
	if err == nil {
		i.To(out, "log file")
	}

	// Change format function to print mock time
	i.Format = func(l *Log, logMsg ...interface{}) string {
		fmtLog := []interface{}{mockTime, l.Flag, fmt.Sprint(logMsg...)}
		log := fmt.Sprintf("%v [%v] - %v\n", fmtLog...)
		return log
	}
	// --------------------------------------------------------------------

	// Mute
	i.Mute()
	// Log hello world
	// Shouldn't be logged
	i.Logf("%v is of type %T", "Hello world", "Hello world")
	// Unmute
	i.Unmute()
	// Log Bye world
	// Output: 2020-01-01 00:00:00 +0000 UTC [INFO] - 'Bye world' is of type string
	i.Logf("%v is of type %T", "'Bye world'", "Bye world")

	// Check the log file
	// --------------------------------------------------------------------
	f, err := os.Open(logFile)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	// Store the line of the log file
	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err != nil {
		t.Fatal(err)
	}

	// Mute not working, more or less than 1 log writted
	if len(lines) != 1 {
		t.Fatalf("number of line expected : 1 \nnumber of line found : %v", len(lines))
	}

	// Check the log format
	output := "2020-01-01 00:00:00 +0000 UTC [INFO] - 'Bye world' is of type string"
	if lines[0] != output {
		t.Fatalf("Expected output : \n%v \ngot : %v", output, lines[0])
	}
	// --------------------------------------------------------------------
}
