package log

import (
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

type Metadata struct {
	Name  string
	Value string
}

type FmtOpts struct {
	Context    []Metadata
	LogLvlName string
	Metadata   []Metadata
	Msg        string
	Timestamp  time.Time
}
type Fmt func(FmtOpts) string

type Std func(string)

type Log struct {
	Context      []Metadata     // Will be prepended to metadata on all log entries
	MinLogLvl    LogLvl         // Minimal log level to output
	Fmt          Fmt            // Log message formatter
	Stderr       Std            // Log message outputter for Debug, Verbose and Info
	Stdout       Std            // Log message outputter for Warning and Error
	TimeLocation *time.Location // Timestamp location/time zone setting
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
	output += " [" + opts.LogLvlName + "] " + opts.Msg

	for i := 0; i < len(opts.Metadata); i++ {
		output += " " + opts.Metadata[i].Name + ": " + opts.Metadata[i].Value
	}

	return output + "\n"
}

func DefaultStderr(msg string) {
	os.Stderr.WriteString(msg)
}

func DefaultStdout(msg string) {
	os.Stdout.WriteString(msg)
}

func (log *Log) SetDefaultValues() {
	if log.MinLogLvl == 0 {
		log.MinLogLvl = 3
	}
	log.Fmt = DefaultFmt
	log.Stderr = DefaultStderr
	log.Stdout = DefaultStdout
	log.TimeLocation, _ = time.LoadLocation("UTC")
}

func (log *Log) Error(msg string) {
	if log.MinLogLvl >= Error {
		log.Stderr(log.Fmt(FmtOpts{
			Timestamp:  time.Now().In(log.TimeLocation),
			LogLvlName: LogNameShort(Error),
			Msg:        msg,
			Metadata:   log.Context,
		}))
	}
}
func (log *Log) ErrorM(msg string, metadata []Metadata) {
	if log.MinLogLvl >= Error {
		log.Stderr(log.Fmt(FmtOpts{
			Timestamp:  time.Now().In(log.TimeLocation),
			LogLvlName: LogNameShort(Error),
			Msg:        msg,
			Metadata:   append(log.Context, metadata[:]...),
		}))
	}
}

func (log *Log) Warn(msg string) {
	if log.MinLogLvl >= Warn {
		log.Stderr(log.Fmt(FmtOpts{
			Timestamp:  time.Now().In(log.TimeLocation),
			LogLvlName: LogNameShort(Warn),
			Msg:        msg,
			Metadata:   log.Context,
		}))
	}
}
func (log *Log) WarnM(msg string, metadata []Metadata) {
	if log.MinLogLvl >= Warn {
		log.Stderr(log.Fmt(FmtOpts{
			Timestamp:  time.Now().In(log.TimeLocation),
			LogLvlName: LogNameShort(Warn),
			Msg:        msg,
			Metadata:   append(log.Context, metadata[:]...),
		}))
	}
}

func (log *Log) Info(msg string) {
	if log.MinLogLvl >= Info {
		log.Stdout(log.Fmt(FmtOpts{
			Timestamp:  time.Now().In(log.TimeLocation),
			LogLvlName: LogNameShort(Info),
			Msg:        msg,
			Metadata:   log.Context,
		}))
	}
}
func (log *Log) InfoM(msg string, metadata []Metadata) {
	if log.MinLogLvl >= Info {
		log.Stdout(log.Fmt(FmtOpts{
			Timestamp:  time.Now().In(log.TimeLocation),
			LogLvlName: LogNameShort(Info),
			Msg:        msg,
			Metadata:   append(log.Context, metadata[:]...),
		}))
	}
}

func (log *Log) Verbose(msg string) {
	if log.MinLogLvl >= Verbose {
		log.Stdout(log.Fmt(FmtOpts{
			Timestamp:  time.Now().In(log.TimeLocation),
			LogLvlName: LogNameShort(Verbose),
			Msg:        msg,
			Metadata:   log.Context,
		}))
	}
}
func (log *Log) VerboseM(msg string, metadata []Metadata) {
	if log.MinLogLvl >= Verbose {
		log.Stdout(log.Fmt(FmtOpts{
			Timestamp:  time.Now().In(log.TimeLocation),
			LogLvlName: LogNameShort(Verbose),
			Msg:        msg,
			Metadata:   append(log.Context, metadata[:]...),
		}))
	}
}

func (log *Log) Debug(msg string) {
	if log.MinLogLvl >= Debug {
		log.Stdout(log.Fmt(FmtOpts{
			Timestamp:  time.Now().In(log.TimeLocation),
			LogLvlName: LogNameShort(Debug),
			Msg:        msg,
			Metadata:   log.Context,
		}))
	}
}
func (log *Log) DebugM(msg string, metadata []Metadata) {
	if log.MinLogLvl >= Debug {
		log.Stdout(log.Fmt(FmtOpts{
			Timestamp:  time.Now().In(log.TimeLocation),
			LogLvlName: LogNameShort(Debug),
			Msg:        msg,
			Metadata:   append(log.Context, metadata[:]...),
		}))
	}
}
