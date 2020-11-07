package resolver

import graphql "github.com/graph-gophers/graphql-go"

type BranchResolver struct {
	id         graphql.ID
	name       string
	repoName   string
	commitHash string
}

func (r *BranchResolver) Id() graphql.ID {
	return r.id
}

func (r *BranchResolver) Name() string {
	return r.name
}

func (r *BranchResolver) Commit() *CommitResolver {
	return &CommitResolver{
		id:       graphql.ID(r.repoName + " " + r.commitHash),
		repoName: r.repoName,
		hash:     r.commitHash,
	}
}
