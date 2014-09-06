package vagrant

import (
	"os"
	"os/user"
)

func homeDir() string {
	usr, err := user.Current()
	var homeDir string
	if err == nil {
		homeDir = usr.HomeDir
	} else {
		// Maybe it's cross compilation without cgo support. (darwin, unix)
		homeDir = os.Getenv("HOME")
	}
	return homeDir
}

func getLongLen(i1, i2 int) int {
	if i1 > i2 {
		return i1
	} else if i1 < i2 {
		return i2
	} else {
		return i2
	}
}
