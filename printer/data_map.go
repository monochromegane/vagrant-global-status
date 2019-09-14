package printer

import (
	"github.com/monochromegane/vagrant-global-status"
)

// ToShort returns the first 7 characters of the id.
func getDataMap(id *vagrant.Id, m *vagrant.Machine) map[string]interface{} {
	data := map[string]interface{}{
		"LocalDataPath":   m.LocalDataPath,
		"Name":            m.Name,
		"Provider":        m.Provider,
		"State":           m.State,
		"VagrantfileName": m.VagrantfileName,
		"VagrantfilePath": m.VagrantfilePath,
		"UpdatedAt":       m.UpdatedAt,
		"ExtraData":       m.ExtraData,
	}

	if id != nil {
		data["Id"] = string(*id)
		data["ShortId"] = id.ToShort()
	}

	return data
}
