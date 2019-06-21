package config

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

type(
	Config struct {
		Server Server
	}

	Server struct {
		Host string
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
