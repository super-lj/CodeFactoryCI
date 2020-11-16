package main

import (
	"ci/method"
	code_factory_ci "ci/thrift_gen/code/factory/ci"
	"ci/util"
	"code.byted.org/gopkg/logs"
	"context"
)

func IsTargetRepoUpdated(ctx context.Context, r *code_factory_ci.IsTargetRepoUpdatedRequest) (*code_factory_ci.IsTargetRepoUpdatedResponse, error) {
	resp, err := method.IsTargetRepoUpdated(ctx, r)
	logs.Infof("IsTargetRepoUpdated request: %s resp: %v", util.MarshallOrElseEmpty(r), util.MarshallOrElseEmpty(resp))
	return resp, err
}

func FetchTargetRepoLastCommit(ctx context.Context, r *code_factory_ci.FetchTargetRepoLastCommitRequest) (*code_factory_ci.FetchTargetRepoLastCommitResonse, error) {
	resp, err := method.FetchTargetRepoLastCommit(ctx, r)
	logs.Infof("FetchTargetRepoLastCommit request: %s resp: %v", util.MarshallOrElseEmpty(r), util.MarshallOrElseEmpty(resp))
	return resp, err
}
