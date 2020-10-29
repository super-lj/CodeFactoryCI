package resolver

import graphql "github.com/graph-gophers/graphql-go"

type BranchResolver struct {
	name     string
	repoName string
	commitID graphql.ID
}

func (r *BranchResolver) Name() string {
	return r.name
}

func (r *BranchResolver) Commit() *CommitResolver {
	return &CommitResolver{
		repoName: r.repoName,
		id:       r.commitID,
	}
}
