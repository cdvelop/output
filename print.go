package output

import "fmt"

// "github.com/fatih/color"

// opt: ok,err,warn,info
func printNEW(opt, message string, a ...interface{}) {

	// switch opt {
	// case "ok":
	// 	color.Green(message, a...)
	// case "err":
	// 	color.Red(message, a...)
	// case "warn":
	// 	color.Yellow(message, a...)
	// case "info":
	// 	color.Cyan(message, a...)
	// default:
	// 	color.White(message, a...)
	// }

	// color.Unset()
}

// options: ok,err,warn,info
func print(message string, options ...string) {
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

	fmt.Printf("\033[%sm%s\033[0m", color, message)
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
