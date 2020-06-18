package config

type (
	// Config ...
	Config struct {
		Server   ServerConfig   `yaml:"server"`
		Database DatabaseConfig `yaml:"database"`
		API      APIConfig      `yaml:"api"`
	}

	// ServerConfig ...
	ServerConfig struct {
		Port string `yaml:"port"`
	}

	// DatabaseConfig ...
	DatabaseConfig struct {
		Master string `yaml:"master"`
	}

	// APIConfig ...
	APIConfig struct {
		MP     string `yaml:"mp"`
		Outlet string `yaml:"outlet"`
		DO     string `yaml:"do"`
	}
)
