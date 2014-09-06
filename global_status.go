package vagrant

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

func GlobalStatus() ([]string, error) {
	index, err := readMachineIndex()
	if err != nil {
		return []string{}, err
	}
	machineIndex := NewMachineIndex(index)
	return machineIndex.GlobalStatuses(), nil
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
