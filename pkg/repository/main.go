package repository

type Repos struct {
	Fs FsRepositoryInterface
	Prompt PromptInterface
}

func NewRepos() Repos {
	return Repos{
		Fs: &FsRepository{},
		Prompt: &Prompt{},
	}
}

func NewMockRepos() Repos {
	return Repos{
		Fs: &FsMockRepository{},
		Prompt: &Prompt{},
	}
}
