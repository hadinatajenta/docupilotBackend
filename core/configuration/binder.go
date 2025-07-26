package configuration

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func LoadConfig(path string) (Config, error) {
	var cfg Config

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(data, &cfg)
	return cfg, err
}
