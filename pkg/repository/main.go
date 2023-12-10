package repository

type Repos struct {
	Fs     FsRepositoryInterface
}

func NewRepos() Repos {
	return Repos{
		Fs:     &FsRepository{},
	}
}

// func NewMockRepos(fsmock FsMockRepository) Repos {
// 	return Repos{
// 		Fs:     &fsmock,
// 	}
// }
