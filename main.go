// package cdep is the Cdep library.
package cdep

import "os"

// Cdep context.
type Cdep struct {
	Home
	conf *Config
}

// New creates new Cdep context with default configuration.
func New() *Cdep {
	return &Cdep{
		conf: defaultConfig(),
	}
}

// New creates new Cdep context with specific configuration.
func NewWith(cfg *Config) *Cdep {
	return &Cdep{
		conf: cfg,
	}
}

func (c *Cdep) Init() {
	c.homepath = c.conf.homePath

	if err := os.MkdirAll(c.homepath, os.ModePerm); err != nil {
		panic(err)
	}
}
