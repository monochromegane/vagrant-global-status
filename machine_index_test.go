package vagrant

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGlobalStatuses(t *testing.T) {

	index := `
	{
		"version": 1,
		"machines": {
			"a0123456789012345678901234567890": {
				"local_data_path": "/Users/user-a/.vagrant",
				"name": "default",
				"provider": "virtualbox",
				"state": "running",
				"vagrantfile_name": null,
				"vagrantfile_path": "/Users/user-a/vagrant",
				"updated_at": null,
				"extra_data": {
					"box": {
						"name": "centos",
						"provider": "virtualbox",
						"version": "0"
					}
				}
			},
			"b0123456789012345678901234567890": {
				"local_data_path": "/Users/user-b/.vagrant",
				"name": "development",
				"provider": "virtualbox",
				"state": "poweroff",
				"vagrantfile_name": null,
				"vagrantfile_path": "/Users/user-b/vagrant",
				"updated_at": null,
				"extra_data": {
					"box": {
						"name": "centos",
						"provider": "virtualbox",
						"version": "0"
					}
				}
			}
		}
	}
	`

	mi := NewMachineIndex([]byte(index))

	if len(mi.Machines) != 2 {
		t.Errorf("It should equal to %d\n", 2)
	}

	expect := &MachineIndex{
		Version: 1,
		Machines: map[Id]Machine{
			"a0123456789012345678901234567890": {
				LocalDataPath: "/Users/user-a/.vagrant",
				Name: "default",
				Provider: "virtualbox",
				State: "running",
				VagrantfileName: "",
				VagrantfilePath: "/Users/user-a/vagrant",
				UpdatedAt: "",
				ExtraData: ExtraData{
					Box: Box{
						Name: "centos",
						Version: "0",
						Provider: "virtualbox",
					},
				},
			},
			"b0123456789012345678901234567890": {
				LocalDataPath: "/Users/user-b/.vagrant",
				Name: "development",
				Provider: "virtualbox",
				State: "poweroff",
				VagrantfileName: "",
				VagrantfilePath: "/Users/user-b/vagrant",
				UpdatedAt: "",
				ExtraData: ExtraData{
					Box: Box{
						Name: "centos",
						Version: "0",
						Provider: "virtualbox",
					},
				},
			},
		},
	}
	if !reflect.DeepEqual(mi, expect) {
		jsonActual, err := json.Marshal(mi)
		if err != nil {
			t.Fatal(err)
		}

		jsonExpected, err := json.Marshal(expect)
		if err != nil {
			t.Fatal(err)
		}

		t.Log("Failed asserting machine index is correct")
		t.Logf("Expected %s", jsonExpected)
		t.Logf("Got      %s", jsonActual)
	}
}
