package repository

import "path/filepath"

type Repos struct {
	WorkspacePath string
	Fs            FsRepositoryInterface
	Prompt        PromptInterface
	Log           LogRepositoryInterface
}

func New() Repos {
	repos := Repos{
		WorkspacePath: "walkhttp.json",
		Fs:            &FsRepository{},
		Prompt:        &Prompt{},
		Log:           &LogRepository{},
	}

	homedir, err := repos.Fs.HomeDir()
	if err != nil {
		return repos
	}
	repos.WorkspacePath = filepath.Join(homedir, ".walkhttp", "walkhttp.json")

	return repos
}
