package printer

import (
	"github.com/monochromegane/vagrant-global-status"
	"io"
	"text/tabwriter"
	"text/template"
)

func PrintTable(output io.Writer, mi *vagrant.MachineIndex, tmpl string) error {
	t := template.New("table")
	t, err := t.Parse(tmpl)
	if err != nil {
		return err
	}
	w := tabwriter.NewWriter(output, 8, 8, 1, ' ', 0)

	var data map[string]interface{}
	for id, m := range mi.Machines {
		data = getDataMap(&id, &m)
		err = t.Execute(w, data)
		if err != nil {
			return err
		}
	}

	return w.Flush()
}
