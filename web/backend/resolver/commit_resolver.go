package resolver

import (
	"context"
	"fmt"
	"web-backend/mock"

	"github.com/graph-gophers/dataloader"
	graphql "github.com/graph-gophers/graphql-go"
)

type CommitResolver struct {
	repoName string
	id       graphql.ID
}

func (r *CommitResolver) Id() graphql.ID {
	return r.id
}

func (r *CommitResolver) RunsConnection(args struct {
	First *int
	After *graphql.ID
}) (*CommitRunsConnectionResolver, error) {
	// fetch commit info
	fetchKey := fmt.Sprintf("%s,%s", r.repoName, r.id)
	res, err := CommitInfoloader.Load(context.TODO(), dataloader.StringKey(fetchKey))()
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}

	// parse retrieved info
	runs := *res.(*[]mock.MockRun)
	for i, run := range runs {
		if args.After == nil || run.ID == *args.After {
			end := len(runs)
			if args.First != nil {
				end = i + *args.First
			}
			if end > len(runs) {
				end = len(runs)
			}
			commitRunsRx := &CommitRunsConnectionResolver{
				pageInfo: &PageInfoResolver{end != len(runs)},
			}
			for _, r := range runs[i:end] {
				commitRunsRx.edges = append(commitRunsRx.edges, &CommitRunsEdgeResolver{
					cursor: r.ID,
					node: &RunResolver{
						id:             r.ID,
						startTimestamp: r.StartTimestamp,
						duration:       r.Duration,
						status:         r.Status,
						log:            r.Log,
					},
				})
			}
			return commitRunsRx, nil
		}
	}

	return nil, nil
}
