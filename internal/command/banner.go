package command

import "github.com/enuesaa/walkhttp/internal/repository"

func printBanner(repos repository.Repos) {
	// print planning like fiber v2 message
	repos.Log.Printf("┌─────────────────────────────────────────────────────────────────\n")
	repos.Log.Printf("│ walkhttp\n")
	repos.Log.Printf("│\n")
	repos.Log.Printf("│ --origin %s \n", repos.Env.WALKHTTP_URL_BASE)
	repos.Log.Printf("│ --port   3000 \n")
	repos.Log.Printf("│\n")
	repos.Log.Printf("│ Web console: http://localhost:3000/_\n")
	repos.Log.Printf("└─────────────────────────────────────────────────────────────────\n")
	repos.Log.Printf("\n")
	repos.Log.Printf(" Try `curl http://localhost:3000/` and open the web console.\n")
}
