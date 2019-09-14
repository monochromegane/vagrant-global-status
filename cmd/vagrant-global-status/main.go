package main

import (
	"flag"
	"fmt"
	"github.com/monochromegane/vagrant-global-status"
	"github.com/monochromegane/vagrant-global-status/printer"
	"os"
)

func main() {
	format := flag.String(
		"format",
		"table",
		"The format to return the details in. Can be table, text, json.",
	)
	tmpl := flag.String(
		"template",
		"{{.ShortId}}\t{{.Name}}\t{{.Provider}}\t{{.State}}\t{{.VagrantfilePath}}",
		"The template to output. When using format is table, tab characters specify columns.\nThis option is ignored when format is json.",
	)
	flag.Parse()

	machineIndex, err := vagrant.GetMachineIndex()
	if err == nil {
		switch *format {
		case "json":
			err = printer.PrintJson(os.Stdout, machineIndex)
		case "table":
			err = printer.PrintTable(os.Stdout, machineIndex, *tmpl)
		case "text":
			err = printer.PrintText(os.Stdout, machineIndex, *tmpl)
		default:
			err = fmt.Errorf("Unknown output format %s", *format)
		}
	}

	if err != nil {
		fmt.Printf("Unexpected error %v", err)
		os.Exit(1)
	}
}
