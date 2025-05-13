package repository

type Repos struct {
	Fs     FsRepositoryInterface
	Prompt PromptInterface
	Log    LogRepositoryInterface
	DB     DBRepositoryInterface
	Env    Env
}

func New() (Repos, error) {
	env, err := NewEnv()
	if err != nil {
		return Repos{}, err
	}

	repos := Repos{
		Fs:     &FsRepository{},
		Prompt: &Prompt{},
		Log:    &LogRepository{},
		DB: &DBRepository{
			data: map[string]interface{}{},
		},
		Env: env,
	}
	return repos, nil
}
