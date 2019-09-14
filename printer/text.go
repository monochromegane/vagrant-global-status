package printer

import (
	"github.com/monochromegane/vagrant-global-status"
	"io"
	"text/template"
)

func PrintText(output io.Writer, mi *vagrant.MachineIndex, tmpl string) error {
	t := template.New("text")
	t, _ = t.Parse(tmpl)

	var data map[string]interface{}
	for id, m := range mi.Machines {
		data = getDataMap(&id, &m)
		err := t.Execute(output, data)
		if err != nil {
			return err
		}
	}

	return nil
}
