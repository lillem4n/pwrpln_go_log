package go_log

import (
	"fmt"
	"os"
	"time"
)

type LogLvl byte

const (
	Error   LogLvl = 1
	Warn    LogLvl = 2
	Info    LogLvl = 3
	Verbose LogLvl = 4
	Debug   LogLvl = 5
)

type FmtOpts struct {
	Context    []interface{}
	LogLvlName string
	Parts      []interface{}
	Timestamp  time.Time
}
type Fmt func(FmtOpts) string

type Std func(string)

type Log struct {
	Context      []interface{}  // Will be prepended to metadata on all log entries
	MinLogLvl    LogLvl         // Minimal log level to output
	Fmt          Fmt            // Log message formatter
	Stderr       Std            // Log message outputter for Debug, Verbose and Info
	Stdout       Std            // Log message outputter for Warning and Error
	TimeLocation *time.Location // Timestamp location/time zone setting
}

func LogLvlFromStr(logLvl string) LogLvl {
	if logLvl == "Error" {
		return Error
	} else if logLvl == "Warn" {
		return Warn
	} else if logLvl == "Info" {
		return Info
	} else if logLvl == "Verbose" {
		return Verbose
	} else if logLvl == "Debug" {
		return Debug
	} else {
		return 0
	}
}

func LogName(logLvl LogLvl) string {
	if logLvl == 0 {
		return "Error"
	}
	if logLvl == 1 {
		return "Warn"
	}
	if logLvl == 2 {
		return "Info"
	}
	if logLvl == 3 {
		return "Verbose"
	}
	if logLvl == 4 {
		return "Debug"
	}
	return "Invalid"
}

func LogNameShort(logLvl LogLvl) string {
	if logLvl == 1 {
		return "\033[31mErr\033[0m"
	}
	if logLvl == 2 {
		return "\033[33mWrn\033[0m"
	}
	if logLvl == 3 {
		return "\033[32mInf\033[0m"
	}
	if logLvl == 4 {
		return "\033[34mVer\033[0m"
	}
	if logLvl == 5 {
		return "\033[37mDeb\033[0m"
	}
	return "Invalid"
}

func DefaultFmt(opts FmtOpts) string {
	output := opts.Timestamp.Format("2006-01-02 15:04:05")
	output += " [" + opts.LogLvlName + "] " + fmt.Sprintf("%v", opts.Parts[0])

	for i := 0; i < len(opts.Context); i = i + 2 {
		output += " " + fmt.Sprintf("%v", opts.Context[i]) + ": " + fmt.Sprintf("%v", opts.Context[i+1])
	}
	for i := 1; i < len(opts.Parts); i = i + 2 {
		output += " " + fmt.Sprintf("%v", opts.Parts[i]) + ": " + fmt.Sprintf("%v", opts.Parts[i+1])
	}

	return output + "\n"
}

func DefaultStderr(msg string) {
	os.Stderr.WriteString(msg)
}

func DefaultStdout(msg string) {
	os.Stdout.WriteString(msg)
}

func GetLog() Log {
	log := Log{
		Fmt:       DefaultFmt,
		MinLogLvl: Info,
		Stderr:    DefaultStderr,
		Stdout:    DefaultStdout,
	}
	log.TimeLocation, _ = time.LoadLocation("UTC")

	return log
}

func (log *Log) Error(parts ...interface{}) {
	if log.MinLogLvl >= Error {
		log.Stderr(log.Fmt(FmtOpts{
			Context:    log.Context,
			LogLvlName: LogNameShort(Error),
			Parts:      parts,
			Timestamp:  time.Now().In(log.TimeLocation),
		}))
	}
}

func (log *Log) Warn(parts ...interface{}) {
	if log.MinLogLvl >= Warn {
		log.Stderr(log.Fmt(FmtOpts{
			Context:    log.Context,
			Timestamp:  time.Now().In(log.TimeLocation),
			LogLvlName: LogNameShort(Warn),
			Parts:      parts,
		}))
	}
}

func (log *Log) Info(parts ...interface{}) {
	if log.MinLogLvl >= Info {
		log.Stdout(log.Fmt(FmtOpts{
			Context:    log.Context,
			Timestamp:  time.Now().In(log.TimeLocation),
			LogLvlName: LogNameShort(Info),
			Parts:      parts,
		}))
	}
}

func (log *Log) Verbose(parts ...interface{}) {
	if log.MinLogLvl >= Verbose {
		log.Stdout(log.Fmt(FmtOpts{
			Context:    log.Context,
			Timestamp:  time.Now().In(log.TimeLocation),
			LogLvlName: LogNameShort(Verbose),
			Parts:      parts,
		}))
	}
}

func (log *Log) Debug(parts ...interface{}) {
	if log.MinLogLvl >= Debug {
		log.Stdout(log.Fmt(FmtOpts{
			Context:    log.Context,
			Timestamp:  time.Now().In(log.TimeLocation),
			LogLvlName: LogNameShort(Debug),
			Parts:      parts,
		}))
	}
}
