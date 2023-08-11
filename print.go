package output

import (
	"fmt"
	"os"
)

var format_color string

func init() {
	for _, arg := range os.Args {
		if arg == "dev" {
			format_color = "\033[%sm%s\033[0m"
		}
	}
}

// options: ok,err,warn,info
func print(message string, options ...string) {
	if format_color == "" {
		fmt.Print(message)
	} else {

		var color string

		for _, opt := range options {
			switch opt {
			case "ok":
				color = "32" //green
			case "err":
				color = "31" //red
			case "warn":
				color = "33" //yellow
			case "info":
				color = "36" //magenta blue=34
			default:
				color = "0"
			}
		}

		fmt.Printf(format_color, color, message)
	}
}

func PrintOK(message string) {
	print(message, "ok")
}

func PrintWarning(message string) {
	print(message, "warn")
}
func PrintError(message string) {
	print(message, "err")
}

func PrintInfo(message string) {
	print(message, "info")
}
