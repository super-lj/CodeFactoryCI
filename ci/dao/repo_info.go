package dao

import (
	"ci/client"
	"ci/domain"
	"ci/enumeration"
	"ci/util"
	"strconv"
)

func GetRepoInfoById(id int) (*domain.Repo, util.OpResult) {
	var repo *domain.Repo
	client.DBClient.First(repo, "id = ?", strconv.Itoa(id))
	return repo, util.NewSucOpResult()
}

func GetRepoInfoByStatus(status enumeration.RepoStatusEnum) (*domain.Repo, util.OpResult) {
	var repo *domain.Repo
	client.DBClient.First(repo, "status = ?", strconv.Itoa(int(status)))
	return repo, util.NewSucOpResult()
}
