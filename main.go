// package cdep is the Cdep library.
package cdep

import (
	"os"
	"path/filepath"
)

// Cdep context.
type Cdep struct {
	Home
	conf *Config
}

// New creates new Cdep context with default configuration.
func New() *Cdep {
	return NewWith(defaultConfig())
}

// New creates new Cdep context with specific configuration.
func NewWith(cfg *Config) *Cdep {
	ret := &Cdep{
		conf: cfg,
	}
	ret.init()
	return ret
}

func (c *Cdep) init() {
	c.homepath = c.conf.homePath

	if err := os.MkdirAll(c.homepath, os.ModePerm); err != nil {
		panic(err)
	}
}

// InitProject inits a Cdep project with given path.
func (c *Cdep) InitProject(path string) error {
	cmakeModPath := filepath.Join(path, "lib_cdep")
	if err := os.MkdirAll(cmakeModPath, os.ModePerm); err != nil {
		Error("%v", err)
		return err
	}
	return nil
}
