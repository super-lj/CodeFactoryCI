package dao

import (
	"ci/domain"
	"ci/enumeration"
	"ci/manager"
	"ci/util"
	"code.byted.org/gopkg/gorm"
	"code.byted.org/gopkg/logs"
	"context"
	"fmt"
)

func GetRepoInfoById(ctx context.Context, id int64) (*domain.Repo, util.OpResult) {
	var repo *domain.Repo
	dbr, err := manager.CodeFactoryDBRead.GetConnection()
	if err != nil {
		logs.CtxError(ctx, "[GetRepoInfoById] get db connection err: %v", err)
		return repo, util.NewOpResult(util.ErrDBConnect, "[GetRepoInfoById] get db connection err")
	}
	err = dbr.Where(fmt.Sprintf("repo_id = %d", id)).First(&repo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, util.NewSucOpResult()
		} else {
			return nil, util.NewOpResult(util.ErrDBRead, "[GetRepoInfoById] get db record err")
		}
	}
	return repo, util.NewSucOpResult()
}

func GetRepoInfoByStatus(ctx context.Context, status enumeration.RepoStatusEnum) (*domain.Repo, util.OpResult) {
	var repo *domain.Repo
	dbr, err := manager.CodeFactoryDBRead.GetConnection()
	if err != nil {
		logs.CtxError(ctx, "[GetRepoInfoByStatus] get db connection err: %v", err)
		return repo, util.NewOpResult(util.ErrDBConnect, "[GetRepoInfoById] get db connection err")
	}
	err = dbr.Where(fmt.Sprintf("status = %d", status)).First(&repo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, util.NewSucOpResult()
		} else {
			return nil, util.NewOpResult(util.ErrDBRead, "[GetRepoInfoByStatus] get db record err")
		}
	}
	return repo, util.NewSucOpResult()
}
