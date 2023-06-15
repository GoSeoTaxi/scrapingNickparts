package config

import (
	"flag"
	"github.com/caarlos0/env"
)

type Config struct {
	StartCopy string `env:"START_COPY"`
	Debug     bool   `env:"SERVER_DEBUG"`
	Path      string `env:"TEMP_PATH"`
	URLImport string `env:"URL_IMPORT"`
	URLExport string `env:"URL_EXPORT"`
}

func InitConfig() (*Config, error) {
	var cfg Config

	flag.StringVar(&cfg.StartCopy, "t", "1", "t=1")
	flag.BoolVar(&cfg.Debug, "debug", false, "debug=true")
	flag.StringVar(&cfg.Path, "path", tPath, "path=C:\\temp\\00")
	flag.StringVar(&cfg.URLImport, "urlI", urlI, "urlI=https://avtozzzapchasti.ru/rest/get_items/")
	flag.StringVar(&cfg.URLExport, "urlE", urlE, "urlE=https://avtozzzapchasti.ru/rest/import/")

	flag.Parse()

	err := env.Parse(&cfg)

	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
