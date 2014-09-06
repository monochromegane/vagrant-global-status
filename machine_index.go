package vagrant

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// A MachineIndex stores vagrant's machine-index file contents.
// It returns formatted strings like `vagrant global-status` command output.
type MachineIndex struct {
	Version  int      `json:"version"`
	Machines Machines `json:"machines"`
}

// A Id stores vagrant machine id.
type Id string

// A Machines stores vagrant machines data.
type Machines map[Id]Machine

// A Machine stores vagrant machine data.
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

// A ExtraData stores vagrant extra data.
type ExtraData struct {
	Box Box `json:"box"`
}

// A Box stores vagrant box data.
type Box struct {
	Name     string `json:"name"`
	Provider string `json:"provider"`
	Version  string `json:"version"`
}

// NewMachineIndex returns MachineIndex from bytes of vagrant machine-index contents.
func NewMachineIndex(bytes []byte) MachineIndex {
	var machineIndex MachineIndex
	json.Unmarshal(bytes, &machineIndex)
	return machineIndex
}

// GlobalStatuses returns strings like `vagrant global-status` command output.
func (mi MachineIndex) GlobalStatuses() []string {
	var statuses []string
	for id, m := range mi.Machines {
		statuses = append(statuses,
			fmt.Sprintf(mi.formatString(),
				id.toShort(),
				m.Name,
				m.Provider,
				m.State,
				m.VagrantfilePath,
			))
	}
	return statuses
}

func (mi MachineIndex) formatString() string {
	max := mi.maxColumnLength()
	return fmt.Sprintf("%%s  %%-%ds %%-%ds %%-%ds %%s",
		max["name"],
		max["provider"],
		max["state"],
	)
}

func (mi MachineIndex) maxColumnLength() map[string]int {
	max := map[string]int{
		"name":     0,
		"provider": 0,
		"state":    0,
	}
	for _, m := range mi.Machines {
		max["name"] = getLongLen(max["name"], len(m.Name))
		max["provider"] = getLongLen(max["provider"], len(m.Provider))
		max["state"] = getLongLen(max["state"], len(m.State))
	}
	return max
}

// toShort returns the first 7 characters of the id.
func (id Id) toShort() string {
	reg := regexp.MustCompile("[a-z0-9]{7}")
	return reg.FindString(string(id))
}
