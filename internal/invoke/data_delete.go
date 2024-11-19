package invoke

import (
	"errors"

	"github.com/enuesaa/walkhttp/internal/repository"
)

func (srv *InvokeSrv) Delete(id string) error {
	err := srv.repos.DB.Omit(id)
	if err != nil {
		if errors.Is(err, repository.DBKeyNotFoundError) {
			return nil
		}
		return err
	}
	return nil
}
