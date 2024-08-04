package usecase

import (
	"github.com/enuesaa/walkhttp/internal/repository"
)

func PrintBanner(repos repository.Repos) {
	// print planning like fiber v2 message
	repos.Log.Printf("┌───────────────────────────────────────────────\n")
	repos.Log.Printf("│ walkhttp\n")
	repos.Log.Printf("│\n")
	repos.Log.Printf("│ serving web console on http://localhost:3000\n")
	repos.Log.Printf("└───────────────────────────────────────────────\n")
	repos.Log.Printf("\n")
}
