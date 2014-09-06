package vagrant

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
)

func GlobalStatus() error {

	index, err := readMachineIndex()
	if err != nil {
		return err
	}

	machineIndex := NewMachineIndex(index)

	re := regexp.MustCompile("[a-z0-9]{7}")
	for id, m := range machineIndex.Machines {
		fmt.Printf("%s %s %s %s %s\n", re.FindString(id), m.Name, m.Provider, m.State, m.VagrantfilePath)
	}
	return nil

}

func readMachineIndex() ([]byte, error) {
	index := filepath.Join(
		homeDir(),
		".vagrant.d",
		"data",
		"machine-index",
		"index",
	)
	return ioutil.ReadFile(index)
}

func homeDir() string {
	usr, err := user.Current()
	var homeDir string
	if err == nil {
		homeDir = usr.HomeDir
	} else {
		// Maybe it's cross compilation without cgo support. (darwin, unix)
		homeDir = os.Getenv("HOME")
	}
	return homeDir
}
