package cli

import (
	"encoding/json"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

type WalkinConf struct {
	BaseUrl string `json:"baseUrl"`
}

func CreateInitCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := repos.Fs.CreateDir(".walkin"); err != nil {
				return err
			}

			conf := WalkinConf{
				BaseUrl: "https://",
			}
			if err := repos.Prompt.Ask("BaseUrl", "", &conf.BaseUrl); err != nil {
				return err
			}

			data, err := json.Marshal(conf)
			if err != nil {
				return err
			}
			if err := repos.Fs.Create(".walkin/config.json", data); err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}