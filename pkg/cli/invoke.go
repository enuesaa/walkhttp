package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/enuesaa/walkin/pkg/event"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateInvokeCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invoke",
		Short: "invoke",
		Run: func(cmd *cobra.Command, args []string) {
			url, _ := cmd.Flags().GetString("url")
			res, err := http.Get(url)
			if err != nil {
				log.Fatalf("Error: %s\n", err)
			}
			defer res.Body.Close()
			resbodybytes, err := io.ReadAll(res.Body)
			if err != nil {
				log.Fatalf("Error: %s\n", err)
			}
			fmt.Printf("%s\n", string(resbodybytes))

			evnt := event.Event{
				Method: "GET",
				Url: url,
				Request: event.EventRequest{},
				Response: event.EventResponse{
					Status: res.StatusCode,
				},
			}
			fbyte, err := json.Marshal(evnt)
			if err != nil {
				log.Fatalf("Error: %s\n", err)
			}
			if err := repos.Fs.Create("event.json", fbyte); err != nil {
				log.Fatalf("Error: %s\n", err)
			}
		},
	}
	cmd.Flags().String("url", "https://example.com", "invoke url")

	return cmd
}
