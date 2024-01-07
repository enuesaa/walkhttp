package taskparse

import (
	"github.com/enuesaa/walkin/pkg/repository"
)

func New(repos repository.Repos) TaskParser {
	return TaskParser{
		repos: repos,
	}
}

type TaskParser struct {
	repos repository.Repos
}
