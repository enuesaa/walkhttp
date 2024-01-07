package usecase

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/taskparse"
)

func ParseTaskfile(repos repository.Repos) (taskparse.Taskfile, error) {
	parser := taskparse.New(repos)
	return parser.Read()
}
