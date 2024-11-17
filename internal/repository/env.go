package repository

type Env struct {
	WALKHTTP_URL_BASE string `env:"WALKHTTP_URL_BASE" envDefault:"https://example.com/"`
}
