package repository

import (
	"encoding/json"
	"io"
	"os"

	"github.com/enuesaa/walkhttp/internal/repository/workspace"
)

type WsRepositoryInterface interface {
	Use(path string) error
	Read(path string) (workspace.Workspace, error)
	Write(path string, ws workspace.Workspace) error
}

type WsRepository struct {
	path string
}

func (repo *WsRepository) Use(path string) error {
	repo.path = path

	return nil
}

func (repo *WsRepository) Read(path string) (workspace.Workspace, error) {
	var ws workspace.Workspace
	f, err := os.Open(path)
	if err != nil {
		return ws, err
	}
	defer f.Close()
	fbytes, err := io.ReadAll(f)
	if err != nil {
		return ws, err
	}
	if err := json.Unmarshal(fbytes, &ws); err != nil {
		return ws, err
	}
	return ws, nil
}

func (repo *WsRepository) Write(path string, ws workspace.Workspace) error {
	fbytes, err := json.MarshalIndent(ws, "", "  ")
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Write(fbytes); err != nil {
		return err
	}
	return nil
}
