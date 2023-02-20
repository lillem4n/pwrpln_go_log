# Simple logger for Go

## Installation

`go get -u gitea.larvit.se/pwrpln/go_log`

## Example usage

Most basic usage with default settings:

```go
import "gitea.larvit.se/pwrpln/go_log"

func main() {
	log := go_log.GetLog()
	log.Error("Apocalypse! :O"); // stderr
	log.Warn("The chaos is near"); // stderr
	log.Info("All is well, but this message is important"); // stdout

	// Will not be shown due to default log level-limit is "info"
	log.Verbose("Extra info, likely good in a production environment"); // stdout
	log.Debug("A lot of detailed logs to debug your application"); // stdout
}
```

Set log level:

```go
log := go_log.GetLog()
log.MinLogLvl = go_log.Debug

// Will now show on stdout
log.Debug("A lot of detailed logs to debug your application");
// 2022-10-11 07:13:49 [Deb] A lot of detailed logs to debug your application
```

Using metadata for structured logging:

```go
log.Info("My log msg", "foo", "bar")
// 2022-10-11 07:13:49 [Inf] My log msg foo: bar
```

Setting a logging context to prepend metadata on all log entries:

```go
log := go_log.GetLog()
log.Context = []interface{{"some", "thing"}}

log.Info("A message")
// 2022-10-11 07:13:49 [Inf] A message some: thing

log.Info("Zep", "other", "stuff")
// 2022-10-11 07:13:49 [Inf] A message some: thing other: stuff
```

All available options, and their defaults:

```go
loc, _ := time.LoadLocation("UTC") // See more info at https://pkg.go.dev/time#LoadLocation
log := go_log.Log{
	Context:      []interface{},                // Will be prepended to metadata on all log entries
	MinLogLvl:    go_log.LogLvlFromStr("Info"), // Minimal log level to output
	Fmt:          go_log.DefaultFmt,            // Log message formatter
	Stderr:       go_log.DefaultStderr,         // Log message outputter for Debug, Verbose and Info
	Stdout:       go_log.DefaultStdout,         // Log message outputter for Warning and Error
	TimeLocation: loc,                          // Timestamp location/time zone setting
}
```

Or change them after initialization like this:

```go
log := go_log.GetLog()
log.MinLogLvl = go_log.LogLvlFromStr("Debug")
```