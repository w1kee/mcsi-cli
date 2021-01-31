package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	mcsi "github.com/w1kee/mcsi-lib"
)

func main() {
	args := parseArgs(os.Args)
	if args == nil {
		displayUsage()
		os.Exit(0)
	}

	status, err := mcsi.GetInfo((*args).Url)
	if err != nil {
		color.Red(fmt.Sprintf("error: %s", err.Error()))
		os.Exit(1)
	}
	if status == nil {
		color.Red("error: the json decoding failed, maybe the server didn't send json (??)")
		os.Exit(1)
	}
	if (*status).Online == false {
		color.Red(fmt.Sprintf("error: %s is either not a minecraft server or it's not online", (*status).Hostname))
		os.Exit(1)
	}

	if (*args).All {
		complexDisplay(*status)
	} else {
		simpleDisplay(*status)
	}
}
