package vagrant

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/user"
	"path/filepath"
)

type MachineIndex struct {
	Version  int                `json:"version"`
	Machines map[string]Machine `json:"machines"`
}

type Machine struct {
	LocalDataPath   string    `json:"local_data_path"`
	Name            string    `json:"name"`
	Provider        string    `json:"provider"`
	State           string    `json:"state"`
	VagrantfileName string    `json"vagrantfile_name"`
	VagrantfilePath string    `json:"vagrantfile_path"`
	UpdatedAt       string    `json:"updated_at"`
	ExtraData       ExtraData `json:"extra_data"`
}

type ExtraData struct {
	Box Box `json:"box"`
}

type Box struct {
	Name     string `json:"name"`
	Provider string `json:"provider"`
	Version  string `json:"version"`
}

func GlobalStatus() {
	usr, _ := user.Current()
	homeDir := usr.HomeDir

	index := filepath.Join(homeDir, ".vagrant.d", "data", "machine-index", "index")
	file, err := ioutil.ReadFile(index)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%s\n", string(file))

	var machineIndex MachineIndex
	json.Unmarshal(file, &machineIndex)
	fmt.Printf("Results: %v\n", machineIndex)
}
