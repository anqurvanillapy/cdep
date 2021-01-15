package cdep

// Cdep configuration.
type Config struct {
	homePath string
}

func defaultConfig() *Config {
	return &Config{
		homePath: HomeDir(),
	}
}
