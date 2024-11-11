package repository

type Repos struct {
	Fs     FsRepositoryInterface
	Prompt PromptInterface
	Log    LogRepositoryInterface
	DB     DBRepositoryInterface
}

func New() Repos {
	repos := Repos{
		Fs:     &FsRepository{},
		Prompt: &Prompt{},
		Log:    &LogRepository{},
		DB:     NewDBRepository(),
	}
	return repos
}
