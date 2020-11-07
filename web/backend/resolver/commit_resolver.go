package resolver

import (
	"context"
	"fmt"
	"strconv"
	"web-backend/mock"

	"github.com/graph-gophers/dataloader"
	graphql "github.com/graph-gophers/graphql-go"
)

type CommitResolver struct {
	id       graphql.ID
	repoName string
	hash     string
}

func (r *CommitResolver) Id() graphql.ID {
	return r.id
}

func (r *CommitResolver) Hash() string {
	return r.hash
}

func (r *CommitResolver) Msg() (string, error) {
	// fetch commit info
	fetchKey := fmt.Sprintf("%s,%s", r.repoName, r.hash)
	res, err := CommitInfoloader.Load(context.TODO(), dataloader.StringKey(fetchKey))()
	if err != nil {
		return "", err
	}

	return res.(*mock.CommitInfo).Msg, nil
}

func (r *CommitResolver) RunsConnection(args struct {
	First *int32
	After *graphql.ID
}) (*CommitRunsConnectionResolver, error) {
	// fetch commit info
	fetchKey := fmt.Sprintf("%s,%s", r.repoName, r.hash)
	res, err := CommitInfoloader.Load(context.TODO(), dataloader.StringKey(fetchKey))()
	if err != nil {
		return nil, err
	}
	if res.(*mock.CommitInfo) == nil {
		return nil, nil
	}

	// parse retrieved info
	runs := res.(*mock.CommitInfo).Runs
	for i, run := range runs {
		runId := graphql.ID(r.repoName + " " + r.hash + " " + strconv.Itoa(int(run.Num)))
		if args.After == nil || runId == *args.After {
			end := len(runs)
			if args.First != nil {
				end = i + int(*args.First)
			}
			if end > len(runs) {
				end = len(runs)
			}
			commitRunsRx := &CommitRunsConnectionResolver{
				pageInfo: &PageInfoResolver{end != len(runs)},
			}
			for _, run := range runs[i:end] {
				runId := graphql.ID(r.repoName + " " + r.hash + " " + strconv.Itoa(int(run.Num)))
				commitRunsRx.edges = append(commitRunsRx.edges, &CommitRunsEdgeResolver{
					cursor: runId,
					node: &RunResolver{
						id:             runId,
						num:            run.Num,
						startTimestamp: run.StartTimestamp,
						duration:       run.Duration,
						status:         run.Status,
						log:            run.Log,
					},
				})
			}
			return commitRunsRx, nil
		}
	}

	return nil, nil
}
