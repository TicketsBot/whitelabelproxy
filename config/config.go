package config

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

type(
	Config struct {
		Server Server
		Patreon Patreon
		Database Database
	}

	Server struct {
		Host string
	}

	Patreon struct {
		Secret string
	}

	Database struct {
		Host string
		Port int
		Username string
		Password string
		Database string
		Pool Pool
	}

	Pool struct {
		MaxConnections int
		MaxIdle int
	}
)

var Conf Config

func LoadConfig() {
	raw, err := ioutil.ReadFile("config.toml"); if err != nil {
		panic(err)
	}

	if err = toml.Unmarshal(raw, &Conf); err != nil {
		panic(err)
	}
}
