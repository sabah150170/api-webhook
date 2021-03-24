package config

//configuration
type Configuration struct {
	Server ServerConf
	Environment EnvHook
}

type ServerConf struct {
	Port string
	Host string
}

type EnvHook struct {
	DUMMY_WEBHOOK_URL string
}

//for json
type Names struct {
	Firstname string
	Lastname string
}