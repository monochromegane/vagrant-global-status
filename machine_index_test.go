package vagrant

import "testing"

func TestGlobalStatuses(t *testing.T) {

	index := `
	{
		"version":1,
		"machines":{
			"a0123456789012345678901234567890":{
				"local_data_path":"/Users/user-a/.vagrant",
				"name":"default",
				"provider":"virtualbox",
				"state":"running",
				"vagrantfile_name":null,
				"vagrantfile_path":"/Users/user-a/vagrant",
				"updated_at":null,
				"extra_data":{
					"box":{
						"name":"centos",
						"provider":"virtualbox",
						"version":"0"
					}
				}
			},
			"b0123456789012345678901234567890":{
				"local_data_path":"/Users/user-b/.vagrant",
				"name":"development",
				"provider":"virtualbox",
				"state":"poweroff",
				"vagrantfile_name":null,
				"vagrantfile_path":"/Users/user-b/vagrant",
				"updated_at":null,
				"extra_data":{
					"box":{
						"name":"centos",
						"provider":"virtualbox",
						"version":"0"
					}
				}
			}
		}
	}
	`

	machines := NewMachineIndex([]byte(index))
	statuses := machines.GlobalStatuses()

	if len(statuses) != 2 {
		t.Errorf("It should equal to %d\n", 2)
	}

	expect := []string{
		"a012345  default     virtualbox running  /Users/user-a/vagrant",
		"b012345  development virtualbox poweroff /Users/user-b/vagrant",
	}
	for i, status := range statuses {
		if status != expect[i] {
			t.Errorf("It should equal to %s\n", expect[i])
		}
	}
}
