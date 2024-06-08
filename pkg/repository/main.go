package repository

type Repos struct {
	DB     DBRepositoryInterface
	Fs     FsRepositoryInterface
	Prompt PromptInterface
	Log    LogRepositoryInterface
}

func NewRepos() Repos {
	return Repos{
		DB:     &InmemoryDB{
			data: make(map[string]interface{}),
		},
		Fs:     &FsRepository{},
		Prompt: &Prompt{},
		Log:    &LogRepository{},
	}
}

func NewMockRepos() Repos {
	return Repos{
		DB:     &InmemoryDB{
			data: make(map[string]interface{}),
		},
		Fs:     &FsMockRepository{},
		Prompt: &Prompt{},
		Log:    &LogMockRepository{},
	}
}
