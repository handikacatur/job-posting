package config

type Config struct {
	Env      string         `yaml:"env"`
	API      APIConfig      `yaml:"api"`
	Database DatabaseConfig `yaml:"database"`
}

type APIConfig struct {
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	Driver   string `yaml:"driver"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}
