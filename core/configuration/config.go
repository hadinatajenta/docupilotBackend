package configuration

import "log"

func MustLoad() Config {
	cfg, err := LoadConfig("env/app-config-map.yml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return cfg
}
