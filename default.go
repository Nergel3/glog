package glog

import (
	"fmt"
	"time"

	"github.com/nergel3/glog/colors"
)

/*
 * DefaultFormat is the default format use to log
 */
func DefaultFormat(l *Log, logMsg ...interface{}) string {
	fmtLog := []interface{}{time.Now().UTC(), l.Flag, fmt.Sprint(logMsg...)}
	return fmt.Sprintf("%v [%v] - %v\n", fmtLog...)
}

/*
 * EMERG log the emergency to stdout
 */
var EMERG = Log{
	Flag:   "EMERG",
	Format: DefaultFormat,
	Output: map[string]Output{
		"stdout": NewStdout(colors.Fg.YELLOW),
	},
}

/*
 * ALERT log the alert to stdout
 */
var ALERT = Log{
	Flag:   "ALERT",
	Format: DefaultFormat,
	Output: map[string]Output{
		"stdout": NewStdout(colors.Fg.YELLOW),
	},
}

/*
 * CRIT log the critique to stdout
 */
var CRIT = Log{
	Flag:   "CRIT",
	Format: DefaultFormat,
	Output: map[string]Output{
		"stdout": NewStdout(colors.Fg.YELLOW),
	},
}

/*
 * ERR log the error to stdout
 */
var ERR = Log{
	Flag:   "ERROR",
	Format: DefaultFormat,
	Output: map[string]Output{
		"stdout": NewStdout(colors.Fg.RED),
	},
}

/*
 * WARNING log the warning to stdout
 */
var WARNING = Log{
	Flag:   "WARN",
	Format: DefaultFormat,
	Output: map[string]Output{
		"stdout": NewStdout(colors.Fg.YELLOW),
	},
}

/*
 * NOTICE log the notice to stdout
 */
var NOTICE = Log{
	Flag:   "NOTICE",
	Format: DefaultFormat,
	Output: map[string]Output{
		"stdout": NewStdout(colors.Fg.WHITE),
	},
}

/*
 * INFO log the info type log to stdout
 */
var INFO = Log{
	Flag:   "INFO",
	Format: DefaultFormat,
	Output: map[string]Output{
		"stdout": NewStdout(colors.Fg.YELLOW),
	},
}

/*
 * DEBUG log the debug to stdout
 */
var DEBUG = Log{
	Flag:   "DEBUG",
	Format: DefaultFormat,
	Output: map[string]Output{
		"stdout": NewStdout(colors.Fg.BLUE),
	},
}
