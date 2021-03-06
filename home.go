package cdep

import (
	"os"
	"path/filepath"
)

// Home is the user's home for caching repositories and some configurations.
type Home struct {
	homepath string
}

// HomeDir gets Cdep's user directory.
func HomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(home, ".cdep")
}
