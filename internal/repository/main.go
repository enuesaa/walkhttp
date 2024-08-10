package repository

type Repos struct {
	Ws     WsRepositoryInterface
	Fs     FsRepositoryInterface
	Prompt PromptInterface
	Log    LogRepositoryInterface
}

func New() Repos {
	return Repos{
		Ws: &WsRepository{
			path: "walkhttp.json",
		},
		Fs:     &FsRepository{},
		Prompt: &Prompt{},
		Log:    &LogRepository{},
	}
}
