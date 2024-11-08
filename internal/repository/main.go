package repository

type Repos struct {
	Fs            FsRepositoryInterface
	Prompt        PromptInterface
	Log           LogRepositoryInterface
	DB            DBRepositoryInterface
}

func New() Repos {
	repos := Repos{
		Fs:            &FsRepository{},
		Prompt:        &Prompt{},
		Log:           &LogRepository{},
		DB:            &DBRepository{
			data: map[string]interface{}{},
		},
	}
	return repos
}
