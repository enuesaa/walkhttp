package usecase

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/repository"
)

func CheckConfigFileExists(repos repository.Repos) error {
	if _, err := repos.Conf.Read(); err != nil {
		return fmt.Errorf("config file does not exist")
	}
	return nil
}
