package vagrant

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"path/filepath"
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
	fmt.Printf("%s\n", string(file))
}
