package resolver

import (
	"fmt"
)

type MutationResolver struct{}

func (r *MutationResolver) CreateFile(args struct{ Name string }) *FileinfoResolver {
	fmt.Println(args.Name)

	return &FileinfoResolver{ name: "aa", description: "aa-description" }
}
