package config

type config struct {
	Env envConfig
}

var Config = &config{}

func Initialize() {
	Config.Env.load()
}
