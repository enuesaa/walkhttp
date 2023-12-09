package proxy

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/enuesaa/walkin/internal/repository"
	"github.com/spf13/cobra"
)

func CreateInvokeCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invoke",
		Short: "invoke",
		Run: func(cmd *cobra.Command, args []string) {	
			req, err := http.NewRequest("GET", "https://example.com/", nil)
			if err != nil {
				log.Fatalf("error: %s", err.Error())
			}

			client := http.Client{}
			res, err := client.Do(req)
			if err != nil {
				log.Fatalf("error: %s", err.Error())
			}
			defer res.Body.Close()

			body, _ := io.ReadAll(res.Body)
			fmt.Println(string(body))
		},
	}

	return cmd
}