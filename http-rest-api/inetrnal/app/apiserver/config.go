package apiserver

type Config struct {
	BindAddr string `toml:"bind_addr`
	LogLevel string `toml:"log_level`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":5432",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
