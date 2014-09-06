// Package vagrant implements functions to provide strings like `vagrant global-status` command output.
package vagrant

import (
	"io/ioutil"
	"path/filepath"
)

// GlobalStatus returns strings like `vagrant global-status` command output.
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
