package lvlog

import (
	"io"
	"log"
	"os"
)

// SetOutput sets the output destination for the standard logger.
// It is a wrapper of the log package log.SetOutput func
func SetOutput(w io.Writer) {
	log.SetOutput(w)
}

// SetFlags sets the output flags for the standard logger.
// It is a wrapper of the log package log.SetFlags func
func SetFlags(flag int) {
	log.SetFlags(flag)
}

var levels = NORM

// SetLevel to 'log.ALL' or 'log.NONE' or 'log.INFO | log.ERROR'
func SetLevel(lvls uint8) {
	levels = lvls
}

// representation of all statuses.
const (
	TRACE = uint8(1 << iota)
	DEBUG
	INFO
	WARN
	ERROR
	PANIC
	FATAL
	ALL  = uint8(TRACE | DEBUG | INFO | WARN | ERROR | PANIC | FATAL)
	NORM = uint8(INFO | WARN | ERROR | PANIC | FATAL)
	NONE = uint8(0x0)
)

// status labels. Feel free to change them.
var (
	TraceLabel = "[TRACE]"
	DebugLabel = "[DEBUG]"
	InfoLabel  = " [INFO]"
	WarnLabel  = " [WARN]"
	ErrorLabel = "[ERROR]"
	PanicLabel = "[PANIC]"
	FatalLabel = "[FATAL]"
)

func appendLevel(label string, v []interface{}) []interface{} {
	v = append(v, 0)
	copy((v)[1:], (v)[0:])
	(v)[0] = label
	return v
}

func println(label string, v ...interface{}) {
	v = appendLevel(label, v)
	log.Println(v...)
}

func printf(label, format string, v ...interface{}) {
	format = label + " " + format
	log.Printf(format, v...)
}

func panicln(label string, v ...interface{}) {
	v = appendLevel(label, v)
	log.Panicln(v...)
}

func panicf(label, format string, v ...interface{}) {
	format = label + " " + format
	log.Panicf(format, v...)
}

// Printf is a wrapper of log package log.Printf func
// Is equivalent to lvlog.Infof(), prepends InfoLabel
func Printf(format string, v ...interface{}) {
	if INFO&levels == 0 {
		return
	}
	printf(InfoLabel, format, v...)
}

// Println is a wrapper of log package log.Println func
// Is equivalent to lvlog.Info(), prepends InfoLabel
func Println(v ...interface{}) {
	if INFO&levels == 0 {
		return
	}
	println(InfoLabel, v...)
}

// Print is a wrapper of log package log.Print func
// Is INFO-level, prepends InfoLabel
func Print(v ...interface{}) {
	if INFO&levels == 0 {
		return
	}
	println(InfoLabel, v...)
}

func Trace(v ...interface{}) {
	if TRACE&levels == 0 {
		return
	}
	println(TraceLabel, v...)
}

func Debug(v ...interface{}) {
	if DEBUG&levels == 0 {
		return
	}
	println(DebugLabel, v...)
}

func Info(v ...interface{}) {
	if INFO&levels == 0 {
		return
	}
	println(InfoLabel, v...)
}

func Warn(v ...interface{}) {
	if WARN&levels == 0 {
		return
	}
	println(WarnLabel, v...)
}

func Error(v ...interface{}) {
	if ERROR&levels == 0 {
		return
	}
	println(ErrorLabel, v...)
}

func Panic(v ...interface{}) {
	if PANIC&levels == 0 {
		return
	}
	panicln(PanicLabel, v...)
}

func Fatal(v ...interface{}) {
	if FATAL&levels == 0 {
		return
	}
	println(FatalLabel, v...)
	os.Exit(1)
}

func Tracef(format string, v ...interface{}) {
	if TRACE&levels == 0 {
		return
	}
	printf(TraceLabel, format, v...)
}

func Debugf(format string, v ...interface{}) {
	if DEBUG&levels == 0 {
		return
	}
	printf(DebugLabel, format, v...)
}

func Infof(format string, v ...interface{}) {
	if INFO&levels == 0 {
		return
	}
	printf(InfoLabel, format, v...)
}

func Warnf(format string, v ...interface{}) {
	if WARN&levels == 0 {
		return
	}
	printf(WarnLabel, format, v...)
}

func Errorf(format string, v ...interface{}) {
	if ERROR&levels == 0 {
		return
	}
	printf(ErrorLabel, format, v...)
}

func Panicf(format string, v ...interface{}) {
	if PANIC&levels == 0 {
		return
	}
	panicf(PanicLabel, format, v...)
}

func Fatalf(format string, v ...interface{}) {
	if FATAL&levels == 0 {
		return
	}
	printf(FatalLabel, format, v...)
	os.Exit(1)
}
