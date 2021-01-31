package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
	mcsi "github.com/w1kee/mcsi-lib"
)

var logo string = fmt.Sprintf(`
                    _
 _ __ ___   ___ ___(_)
| '_ ' _ \ / __/ __| |
| | | | | | (__\__ \ |
|_| |_| |_|\___|___/_|

mcsi-cli - version %s

Credits to:
- Anders G. Jørgensen for making https://mcsrvstat.us/ which provides the api (https://spirit55555.dk/)
- w1ke for making the cli (https://w1ke.cz/)

`,
	version,
)

var version string = "1.0"

func displayMOTD(motd []string) {
	for _, line := range motd {
		runeLine := []rune(line)
		fmt.Print("  ")
		for i := 0; i < len(runeLine); i++ {
			if runeLine[i] == '§' {
				i++
				switch runeLine[i] {
				case '0':
					color.Set(color.FgBlack)
				case '1':
					color.Set(color.FgBlue)
				case '2':
					color.Set(color.FgGreen)
				case '3':
					color.Set(color.FgCyan)
				case '4':
					color.Set(color.FgRed)
				case '5':
					color.Set(color.FgMagenta)
				case '6':
					color.Set(color.FgYellow)
				case '7':
					color.Set(color.FgWhite)
				case '8':
					color.Set(color.FgHiBlack)
				case '9':
					color.Set(color.FgHiBlue)
				case 'a':
					color.Set(color.FgHiGreen)
				case 'b':
					color.Set(color.FgHiCyan)
				case 'c':
					color.Set(color.FgHiRed)
				case 'd':
					color.Set(color.FgHiMagenta)
				case 'e':
					color.Set(color.FgHiYellow)
				case 'f':
					color.Set(color.FgHiWhite)
				case 'l':
					color.Set(color.Bold)
				case 'n':
					color.Set(color.Underline)
				case 'o':
					color.Set(color.Italic)
				case 'r':
					color.Unset()
				default: //unsupported formatting code such as §k or §m
				}
			} else {
				fmt.Printf("%c", runeLine[i])
			}
		}
		color.Unset()
		fmt.Print("\n")
	}
}

func simpleDisplay(status mcsi.Status) error {
	//print out basic information
	//motd
	color.Yellow("MOTD:")
	displayMOTD(status.Motd.Raw)
	fmt.Print("\n")

	//players
	color.Yellow(fmt.Sprintf("Players (%d online)\n", status.Players.Online))
	if len(status.Players.List) != 0 {
		for _, player := range status.Players.List {
			fmt.Printf("  %s\n", player)
		}
	}

	return nil
}

func complexDisplay(status mcsi.Status) error {
	niceJSON, err := json.MarshalIndent(status, "", "    ")
	if err != nil {
		return fmt.Errorf("error: complexDisplay: %v", err)
	}

	fmt.Printf(string(niceJSON))

	return nil
}

func displayUsage() {
	fmt.Printf(color.YellowString(logo))
	fmt.Printf(color.YellowString("Usage:\n"))
	fmt.Printf(color.YellowString("%s <url> - gets info and outputs the basics\n"), os.Args[0])
	fmt.Printf(color.YellowString("%s -a <url> - gets info and outputs the whole json, except for the icon\n"), os.Args[0])
	fmt.Printf(color.YellowString("%s -i <url> - !UNIMPLEMENTED! gets the icon and outputs it\n"), os.Args[0])
}
