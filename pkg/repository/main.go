package repository

type Repos struct {
	Fs     FsRepositoryInterface
	Prompt PromptInterface
	Log    LogRepositoryInterface
}

func NewRepos() Repos {
	return Repos{
		Fs:     &FsRepository{},
		Prompt: &Prompt{},
		Log:    &LogRepository{},
	}
}

func NewMockRepos() Repos {
	return Repos{
		Fs:     &FsMockRepository{},
		Prompt: &Prompt{},
		Log:    &LogMockRepository{},
	}
}
