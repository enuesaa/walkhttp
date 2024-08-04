package repository

type Repos struct {
	Ws     WsRepositoryInterface
	DB     DBRepositoryInterface
	Fs     FsRepositoryInterface
	Prompt PromptInterface
	Log    LogRepositoryInterface
}

func New() Repos {
	return Repos{
		Ws: &WsRepository{
			path: "walkhttp.json",
		},
		DB: &InmemoryDB{
			data: make(map[string]interface{}),
		},
		Fs:     &FsRepository{},
		Prompt: &Prompt{},
		Log:    &LogRepository{},
	}
}
