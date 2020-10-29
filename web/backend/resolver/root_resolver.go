package resolver

import "web-backend/mock"

type RootResolver struct{}

func (r *RootResolver) Repos(args struct{ Name *string }) []*RepoResolver {
	var repoRxs []*RepoResolver
	if args.Name != nil {
		repoRxs = append(repoRxs, &RepoResolver{*args.Name})
	} else {
		for _, name := range mock.GetRepoNames() {
			repoRxs = append(repoRxs, &RepoResolver{name})
		}
	}
	return repoRxs
}
