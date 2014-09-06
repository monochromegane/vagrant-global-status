package vagrant

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type MachineIndex struct {
	Version  int      `json:"version"`
	Machines Machines `json:"machines"`
}

type Id string

type Machines map[Id]Machine

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

func NewMachineIndex(bytes []byte) MachineIndex {
	var machineIndex MachineIndex
	json.Unmarshal(bytes, &machineIndex)
	return machineIndex
}

func (mi MachineIndex) GlobalStatuses() []string {
	var statuses []string
	for id, m := range mi.Machines {
		statuses = append(statuses,
			fmt.Sprintf("%s %s %s %s %s",
				id.toShort(),
				m.Name,
				m.Provider,
				m.State,
				m.VagrantfilePath,
			))
	}
	return statuses
}

func (id Id) toShort() string {
	reg := regexp.MustCompile("[a-z0-9]{7}")
	return reg.FindString(string(id))
}
