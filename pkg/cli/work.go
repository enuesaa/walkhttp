package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateWorkCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "work",
		Short: "work",
		Run: func(cmd *cobra.Command, args []string) {
			host, _ := cmd.Flags().GetString("host")

			listres, err := http.Get(fmt.Sprintf("http://%s/files", host))
			if err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
			defer listres.Body.Close()
			fmt.Printf("status(%s): %s", listres.Request.URL, listres.Status)
			fbytes, err := io.ReadAll(listres.Body)
			var files SyncFileList
			if err := json.Unmarshal(fbytes, &files); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}

			for _, file := range files.Items {
				if file.IsDir {
					if err := repos.Fs.CreateDir(file.Path); err != nil {
						log.Fatalf("Error: %s", err.Error())
					}
				} else {
					res, err := http.Get(fmt.Sprintf("http://%s/files/%s", host, file.Path))
					if err != nil {
						log.Fatalf("Error: %s", err.Error())
					}
					defer res.Body.Close()
					fmt.Printf("status(%s): %s\n", res.Request.URL, res.Status)
					fbytes, err := io.ReadAll(res.Body)
					if err != nil {
						log.Fatalf("Error: %s", err.Error())
					}
					if err := repos.Fs.Create(file.Path, fbytes); err != nil {
						log.Fatalf("Error: %s", err.Error())
					}
				}
			}
		},
	}
	cmd.Flags().String("host", "localhost:3001", "host")

	return cmd
}