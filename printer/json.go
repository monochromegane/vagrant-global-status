package printer

import (
	"encoding/json"
	"github.com/monochromegane/vagrant-global-status"
	"io"
)

func PrintJson(output io.Writer, mi *vagrant.MachineIndex) error {
	b, err := json.Marshal(mi)
	if err != nil {
		return err
	}

	_, err = output.Write(b)
	return err
}
