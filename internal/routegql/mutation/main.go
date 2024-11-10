package mutation

import (
	"github.com/enuesaa/walkhttp/internal/repository"
)

type MutationResolver struct {
	Repos repository.Repos
}
