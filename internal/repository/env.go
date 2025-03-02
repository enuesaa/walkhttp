package repository

import "github.com/caarlos0/env/v10"

func NewEnv() (Env, error) {
	e := Env{}
	if err := env.Parse(&e); err != nil {
		return e, err
	}
	return e, nil
}

type Env struct {
	WALKHTTP_URL_BASE      string `env:"WALKHTTP_URL_BASE" envDefault:"https://example.com"`
	WALKHTTP_PERSIST_COUNT int    `env:"WALKHTTP_PERSIST_COUNT" envDefault:"1000"`
	WALKHTTP_REQUEST_BODY  string `env:"WALKHTTP_REQUEST_BODY" envDefault:"include"`
	WALKHTTP_RESPONSE_BODY string `env:"WALKHTTP_RESPONSE_BODY" envDefault:"include"`
}
