package repository

import (
	"github.com/caarlos0/env/v10"
)

type Repos struct {
	Fs     FsRepositoryInterface
	Prompt PromptInterface
	Log    LogRepositoryInterface
	DB     DBRepositoryInterface
	Env    Env
}

func New() (Repos, error) {
	e := Env{}
	if err := env.Parse(&e); err != nil {
		return Repos{}, err
	}

	repos := Repos{
		Fs:     &FsRepository{},
		Prompt: &Prompt{},
		Log:    &LogRepository{},
		DB:     NewDBRepository(),
		Env:    e,
	}
	return repos, nil
}
