package repository

type Repos struct {
	DB     DBRepositoryInterface
	Fs     FsRepositoryInterface
	Prompt PromptInterface
	Log    LogRepositoryInterface
}

func New() Repos {
	return Repos{
		DB: &InmemoryDB{
			data: make(map[string]interface{}),
		},
		Fs:     &FsRepository{},
		Prompt: &Prompt{},
		Log:    &LogRepository{},
	}
}
