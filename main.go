// Package cdep is the Cdep library.
package cdep

import (
	"io/ioutil"
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

// NewWith creates new Cdep context with specific configuration.
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

const projCmakeTmpl = `message(STATUS "Hello, Cdep!")
`

// InitProject inits a Cdep project with given path.
func (c *Cdep) InitProject(path string) error {
	modPath := filepath.Join(path, "lib_cdep")
	if err := os.MkdirAll(modPath, os.ModePerm); err != nil {
		Error("%v", err)
		return err
	}

	modFile := filepath.Join(modPath, "FindCdep.cmake")
	if err := ioutil.WriteFile(modFile, []byte(projCmakeTmpl), 0644); err != nil {
		return err
	}

	return nil
}
