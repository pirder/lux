package main

import (
	"C"
	"fmt"
	"github.com/fatih/color"
	"github.com/iawia002/lux/app"
	"os"
	"strings"
)

func main() {
	if err := app.New().Run(os.Args); err != nil {
		fmt.Fprintf(
			color.Output,
			"Run %s failed: %s\n",
			color.CyanString("%s", app.Name), color.RedString("%v", err),
		)
		os.Exit(1)
	}
}

//export cli
func cli(in *C.char) *C.char {
	var arg = C.GoString(in)
	var inputArgs = strings.Split(arg, ",")
	var args []string
	args = append(args, "lux")
	args = append(args, inputArgs...)
	var errStr = ""
	if err := app.New().Run(args); err != nil {
		errStr = fmt.Sprintf(
			"Run %s failed: %s\n",
			color.CyanString("%s", app.Name), color.RedString("%v", err),
		)
	}
	return C.CString(errStr)
}
