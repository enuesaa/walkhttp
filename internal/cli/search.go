package cli

import (
	"fmt"

	"github.com/enuesaa/walkin/internal/repository"
	"github.com/spf13/cobra"
	"golang.org/x/text/search"

	"golang.org/x/text/language"
)

func CreateSearchCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search",
		Short: "search",
		Run: func(cmd *cobra.Command, args []string) {	
			fmt.Println("a")
			matcher := search.New(language.Japanese)
			fmt.Println(matcher.IndexString("こんにちは", "ん"))
		},
	}

	return cmd
}
