// Package vagrant implements functions to provide strings like `vagrant global-status` command output.
package vagrant

import (
	"io/ioutil"
	"path/filepath"
)

// Returns a MachineIndex pointer that can be queried for a list of Vagrant machines
func GetMachineIndex() (*MachineIndex, error) {
	index, err := readMachineIndex()
	if err != nil {
		return nil, err
	}
	return NewMachineIndex(index)
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
