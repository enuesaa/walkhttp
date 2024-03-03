package conf

import (
	"github.com/enuesaa/walkin/pkg/repository"
)

func CreateWalkinDir(repos repository.Repos) error {
	if err := repos.Fs.CreateDir(".walkin"); err != nil {
		return err
	}
	return nil
}
