package interfaces

import (
	"os"
	"path/filepath"
)

func ConfigDir() (dir string, err error) {
	dir, err = os.UserHomeDir()
	if err != nil {
		return
	}

	dir = filepath.Join(dir, ".o2")

	return
}
