package resolver

import (
	"context"
	"web-backend/mock"

	"github.com/graph-gophers/dataloader"
	graphql "github.com/graph-gophers/graphql-go"
)

type RepoResolver struct {
	name string
}

func (r *RepoResolver) Name() string {
	return r.name
}

func (r *RepoResolver) BranchesConnection(args struct {
	First *int
	After *string
}) (*RepoBranchesConnectionResolver, error) {
	// fetch repo info
	res, err := RepoInfoloader.Load(context.TODO(), dataloader.StringKey(r.name))()
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}

	// parse retrieved info
	info := res.(*mock.RepoInfo)
	for i, br := range info.Branches {
		if args.After == nil || br.Name == *args.After {
			end := len(info.Branches)
			if args.First != nil {
				end = i + *args.First
			}
			if end > len(info.Branches) {
				end = len(info.Branches)
			}
			repoBrRx := &RepoBranchesConnectionResolver{
				pageInfo: &PageInfoResolver{end != len(info.Branches)},
			}
			for _, b := range info.Branches[i:end] {
				repoBrRx.edges = append(repoBrRx.edges, &RepoBranchesEdgeResolver{
					cursor: b.Name,
					node: &BranchResolver{
						name:     b.Name,
						repoName: r.name,
						commitID: b.CommitID,
					},
				})
			}
			return repoBrRx, nil
		}
	}
	return nil, nil
}

func (r *RepoResolver) CommitsConnection(args struct {
	First *int
	After *graphql.ID
}) (*RepoCommitsConnectionResolver, error) {
	// fetch repo info
	res, err := RepoInfoloader.Load(context.TODO(), dataloader.StringKey(r.name))()
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}

	// parse retrieved info
	info := res.(*mock.RepoInfo)
	for i, commitID := range info.CommitIDs {
		if args.After == nil || commitID == *args.After {
			end := len(info.CommitIDs)
			if args.First != nil {
				end = i + *args.First
			}
			if end > len(info.CommitIDs) {
				end = len(info.CommitIDs)
			}
			repoCmRx := &RepoCommitsConnectionResolver{
				pageInfo: &PageInfoResolver{end != len(info.Branches)},
			}
			for _, commitID := range info.CommitIDs[i:end] {
				repoCmRx.edges = append(repoCmRx.edges, &RepoCommitsEdgeResolver{
					cursor: commitID,
					node: &CommitResolver{
						repoName: r.name,
						id:       commitID,
					},
				})
			}
			return repoCmRx, nil
		}
	}

	return nil, nil
}
