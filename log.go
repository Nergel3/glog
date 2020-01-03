package glog

import (
	"fmt"
)

// Log - Basics log to io.Writer
type Log struct {
	Flag   string
	Format func(*Log, ...interface{}) string
	Output map[string]Output
}

/*
 * Log write the log to all the output registered (except quiet one)
 */
func (l *Log) Log(data ...interface{}) {
	log := []byte(l.Format(l, data...))
	for _, output := range l.Output {
		output.Write(log)
	}
}

/*
 * Logf write the formatted log to all the output registered (except quiet one)
 */
func (l *Log) Logf(format string, data ...interface{}) {
	fmtData := fmt.Sprintf(format, data...)
	l.Log(fmtData)
}

/*
 * Mute all the output of this Logger
 */
func (l *Log) Mute() {
	for _, output := range l.Output {
		output.Mute()
	}
}

/*
 * Unmute all the output of this Logger
 */
func (l *Log) Unmute() {
	for _, output := range l.Output {
		output.Unmute()
	}
}

/*
 * To add an output to the logger
 */
func (l *Log) To(output Output, as string) {
	l.Output[as] = output
}
