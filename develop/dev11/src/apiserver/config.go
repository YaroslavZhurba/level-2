package apiserver

type Config struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

func NewConfig() *Config {
	return &Config{
		Host : "127.0.0.1",
		Port: ":8080",
	}
}
