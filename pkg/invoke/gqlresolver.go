package invoke

import (
	"github.com/enuesaa/walkhttp/pkg/repository"
)

type Resolver struct {
	repos repository.Repos
	baseUrl string
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
