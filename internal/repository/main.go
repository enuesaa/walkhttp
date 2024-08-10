package repository

type Repos struct {
	WorkspacePath string
	Fs     FsRepositoryInterface
	Prompt PromptInterface
	Log    LogRepositoryInterface
}

func New() Repos {
	return Repos{
		WorkspacePath: "walkhttp.json",
		Fs:     &FsRepository{},
		Prompt: &Prompt{},
		Log:    &LogRepository{},
	}
}
