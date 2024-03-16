package repository

type Repos struct {
	Conf   ConfRepositoryInterface
	Fs     FsRepositoryInterface
	Prompt PromptInterface
	Log    LogRepositoryInterface
}

func NewRepos() Repos {
	return Repos{
		Conf:   &ConfRepository{
			fs: FsRepository{},
		},
		Fs:     &FsRepository{},
		Prompt: &Prompt{},
		Log:    &LogRepository{},
	}
}

func NewMockRepos() Repos {
	return Repos{
		Conf:   &ConfRepository{},
		Fs:     &FsMockRepository{},
		Prompt: &Prompt{},
		Log:    &LogMockRepository{},
	}
}
