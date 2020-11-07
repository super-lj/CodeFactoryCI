package resolver

import (
	"context"
	"web-backend/mock"

	"github.com/graph-gophers/dataloader"
	graphql "github.com/graph-gophers/graphql-go"
)

type RepoResolver struct {
	id   graphql.ID
	name string
}

func (r *RepoResolver) Id() graphql.ID {
	return r.id
}

func (r *RepoResolver) Name() string {
	return r.name
}

func (r *RepoResolver) BranchesConnection(args struct {
	First *int32
	After *graphql.ID
}) (*RepoBranchesConnectionResolver, error) {
	// fetch repo info
	res, err := RepoInfoloader.Load(context.TODO(), dataloader.StringKey(r.name))()
	if err != nil {
		return nil, err
	}
	if res.(*mock.RepoInfo) == nil {
		return nil, nil
	}

	// parse retrieved info
	info := res.(*mock.RepoInfo)
	for i, br := range info.Branches {
		brID := graphql.ID(r.name + " " + br.Name)
		if args.After == nil || brID == *args.After {
			end := len(info.Branches)
			if args.First != nil {
				end = i + int(*args.First)
			}
			if end > len(info.Branches) {
				end = len(info.Branches)
			}
			repoBrRx := &RepoBranchesConnectionResolver{
				pageInfo: &PageInfoResolver{end != len(info.Branches)},
			}
			for _, b := range info.Branches[i:end] {
				brID := graphql.ID(r.name + " " + b.Name)
				repoBrRx.edges = append(repoBrRx.edges, &RepoBranchesEdgeResolver{
					cursor: brID,
					node: &BranchResolver{
						id:         brID,
						name:       b.Name,
						repoName:   r.name,
						commitHash: b.CommitHash,
					},
				})
			}
			return repoBrRx, nil
		}
	}
	return nil, nil
}

func (r *RepoResolver) CommitsConnection(args struct {
	First *int32
	After *graphql.ID
}) (*RepoCommitsConnectionResolver, error) {
	// fetch repo info
	res, err := RepoInfoloader.Load(context.TODO(), dataloader.StringKey(r.name))()
	if err != nil {
		return nil, err
	}
	if res.(*mock.RepoInfo) == nil {
		return nil, nil
	}

	// parse retrieved info
	info := res.(*mock.RepoInfo)
	for i, hash := range info.CommitHashs {
		commitID := graphql.ID(info.Name + " " + hash)
		if args.After == nil || commitID == *args.After {
			end := len(info.CommitHashs)
			if args.First != nil {
				end = i + int(*args.First)
			}
			if end > len(info.CommitHashs) {
				end = len(info.CommitHashs)
			}
			repoCmRx := &RepoCommitsConnectionResolver{
				pageInfo: &PageInfoResolver{end != len(info.Branches)},
			}
			for _, hash := range info.CommitHashs[i:end] {
				commitID := graphql.ID(info.Name + " " + hash)
				repoCmRx.edges = append(repoCmRx.edges, &RepoCommitsEdgeResolver{
					cursor: commitID,
					node: &CommitResolver{
						id:       commitID,
						repoName: r.name,
						hash:     hash,
					},
				})
			}
			return repoCmRx, nil
		}
	}

	return nil, nil
}
