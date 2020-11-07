package resolver

import (
	"web-backend/mock"

	"github.com/graph-gophers/graphql-go"
)

type RootResolver struct{}

func (r *RootResolver) Repos(args struct{ Name *string }) []*RepoResolver {
	var repoRxs []*RepoResolver
	if args.Name != nil {
		repoRxs = append(repoRxs, &RepoResolver{
			id:   graphql.ID(*args.Name),
			name: *args.Name,
		})
	} else {
		for _, name := range mock.GetRepoNames() {
			repoRxs = append(repoRxs, &RepoResolver{
				id:   graphql.ID(name),
				name: name,
			})
		}
	}
	return repoRxs
}
