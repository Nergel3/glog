package glog

import (
	"io"
	"os"

	"github.com/nergel3/glog/colors"
)

/*
 * Interface used for logging
 */
type Output interface {
	Write([]byte) (n int, err error)
	Mute()
	Unmute()
}

/*
 * LogOutput is the logging output use by logger
 */
type LogOutput struct {
	Writer io.Writer
	Quiet  bool
	Color  colors.Color
}

/*
 * Write to the log output writer
 */
func (lo *LogOutput) Write(p []byte) (n int, err error) {
	if !lo.Quiet {
		p = lo.Color.Text(p)
		n, err = lo.Writer.Write(p)
	} else {
		n, err = 0, nil
	}

	return
}

/*
 * Mute the output
 */
func (lo *LogOutput) Mute() {
	lo.Quiet = true
}

/*
 * Unmute the output
 */
func (lo *LogOutput) Unmute() {
	lo.Quiet = false
}

/*
 * NewLogOutput return a default log output
 */
func NewOutput() Output {
	return &LogOutput{
		Writer: os.Stdout,
		Quiet:  false,
		Color:  colors.Fg.WHITE,
	}
}

/*
 * NewStdout return a log output to the standard output with the colors
 */
func NewStdout(color colors.Color) Output {
	return &LogOutput{
		Writer: os.Stdout,
		Quiet:  false,
		Color:  color,
	}
}

/*
 * NewStdout return a log output to the standard output with the colors
 */
func NewFileOutput(logFilePath string) (output Output, err error) {
	f, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err == nil {
		output = &LogOutput{
			Writer: f,
			Quiet:  false,
			Color:  colors.Fg.NONE,
		}
	}

	return output, err
}
