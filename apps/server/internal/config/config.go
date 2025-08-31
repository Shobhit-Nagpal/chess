package config

type Config struct {
	env *Env
}

var config = New()

func New() *Config {
	return &Config{
		env: NewEnv(),
	}
}

func GetConfig() *Config {
	return config
}

func (c *Config) GetEnv() *Env {
	return c.env
}
