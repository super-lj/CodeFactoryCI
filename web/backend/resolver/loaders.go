package resolver

import (
	"context"
	"strings"
	"web-backend/mock"

	"github.com/graph-gophers/dataloader"
	graphql "github.com/graph-gophers/graphql-go"
)

// batch functions
var getRepoInfoBatchFn = func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var results []*dataloader.Result
	for _, name := range keys.Keys() {
		results = append(results, &dataloader.Result{Data: mock.GetRepoInfo(name)})
	}
	return results
}

var getCommitInfoBatchFn = func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var results []*dataloader.Result
	for _, k := range keys.Keys() {
		strs := strings.Split(k, ",")
		if len(strs) != 2 {
			return []*dataloader.Result{}
		}
		commit_info := mock.GetCommitInfo(strs[0], graphql.ID(strs[1]))
		results = append(results, &dataloader.Result{Data: commit_info})
	}
	return results
}

// Loaders
var RepoInfoloader = dataloader.NewBatchedLoader(getRepoInfoBatchFn)
var CommitInfoloader = dataloader.NewBatchedLoader(getCommitInfoBatchFn)
