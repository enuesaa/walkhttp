package command

import "github.com/enuesaa/walkhttp/internal/repository"

func printBanner(repos repository.Repos, port int) {
	repos.Log.Printf("┌─────────────────────────────────────────────────────────────────\n")
	repos.Log.Printf("│ walkhttp\n")
	repos.Log.Printf("│\n")
	repos.Log.Printf("│ --origin %s \n", repos.Env.WALKHTTP_URL_BASE)
	repos.Log.Printf("│ --port   %d \n", port)
	repos.Log.Printf("│\n")
	repos.Log.Printf("│ Web console: http://localhost:%d/_\n", port)
	repos.Log.Printf("└─────────────────────────────────────────────────────────────────\n")
	repos.Log.Printf("\n")
	repos.Log.Printf(" Open the web console and try `curl http://localhost:%d/`\n", port)
	repos.Log.Printf("\n")
}
