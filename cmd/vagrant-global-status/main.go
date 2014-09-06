package main

import (
	"fmt"
	"os"

	"github.com/monochromegane/vagrant-global-status"
)

func main() {
	statuses, err := vagrant.GlobalStatus()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	for _, s := range statuses {
		fmt.Println(s)
	}
}
