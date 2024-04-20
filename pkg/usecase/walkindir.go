package usecase

import (
	"path/filepath"

	"github.com/enuesaa/walkin/pkg/repository"
)

func CreateWalkindir(repos repository.Repos) error {
	homedir, err := repos.Fs.HomeDir()
	if err != nil {
		return err
	}
	walkindir := filepath.Join(homedir, ".walkin")
	
	return repos.Fs.CreateDir(walkindir)
}
