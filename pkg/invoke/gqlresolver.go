package invoke

import (
	"github.com/enuesaa/walkin/pkg/repository"
)

type Resolver struct {
	repos repository.Repos
}

type queryResolver struct {
	*Resolver
}

type mutationResolver struct {
	*Resolver
}

type subscriptionResolver struct {
	*Resolver
}
