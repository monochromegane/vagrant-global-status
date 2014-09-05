package vagrant

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"regexp"
)

func GlobalStatus() {
	usr, _ := user.Current()
	homeDir := usr.HomeDir

	index := filepath.Join(homeDir, ".vagrant.d", "data", "machine-index", "index")
	file, err := ioutil.ReadFile(index)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	var machineIndex MachineIndex
	json.Unmarshal(file, &machineIndex)

	re := regexp.MustCompile("[a-z0-9]{7}")
	for id, m := range machineIndex.Machines {
		fmt.Printf("%s %s %s %s %s\n", re.FindString(id), m.Name, m.Provider, m.State, m.VagrantfilePath)
	}

}
