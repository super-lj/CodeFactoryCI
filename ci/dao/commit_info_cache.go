package dao

import (
	"ci/client"
	"ci/domain"
	"ci/util"
	"context"
	"encoding/json"
	"fmt"
)

// GetCommitByIdFromCache returns a commit associated with given commit id
func GetCommitByIdFromCache(ctx context.Context, id int64) (*domain.Commit, util.OpResult) {
	key := util.GetCommitRedisKey(id)
	res, err := client.RedisService.Get(key).Result()
	if err != nil {
		return nil, util.NewOpResult(util.ErrRedis, fmt.Sprintf("get commit from redis err, id: %s", id))
	}
	var commit *domain.Commit
	err = json.Unmarshal([]byte(res), commit)
	if err != nil {
		return nil, util.NewOpResult(util.ErrSystemInternal, fmt.Sprintf("unmarshal from string err"))
	}
	return commit, util.NewSucOpResult()
}
