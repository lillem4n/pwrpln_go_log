# Simple logger for Go

## Installation

`go get -u https://gitea.larvit.se/pwrpln/go_log`

## Example usage

```go
import "gitea.larvit.se/pwrpln/go_log"

func main() {
	log := Log{}
	log.SetDefaultValues{}
	log.Info("My little log message")
}
```