package mock

import (
	graphql "github.com/graph-gophers/graphql-go"
)

type MockRepo struct {
	Name     string
	Branches []MockBranch
	Commits  []MockCommit
}

type MockBranch struct {
	Name     string
	CommitID graphql.ID
}

type MockCommit struct {
	ID   graphql.ID
	Runs []MockRun
}

type MockRun struct {
	ID             graphql.ID
	StartTimestamp int32
	Duration       int32
	Status         string
	Log            string
}

var MockRepoData = []MockRepo{
	{
		Name: "test_repo_a",
		Branches: []MockBranch{
			{
				Name:     "main",
				CommitID: graphql.ID("0000003"),
			},
			{
				Name:     "dev",
				CommitID: graphql.ID("0000002"),
			},
		},
		Commits: []MockCommit{
			{
				ID: graphql.ID("0000001"),
				Runs: []MockRun{
					{
						ID:             graphql.ID("1"),
						StartTimestamp: 0,
						Duration:       100,
						Status:         "SUCCEED",
						Log:            "",
					},
				},
			},
			{
				ID: graphql.ID("0000002"),
				Runs: []MockRun{
					{
						ID:             graphql.ID("2"),
						StartTimestamp: 100,
						Duration:       100,
						Status:         "FAILED",
						Log:            "",
					},
				},
			},
			{
				ID: graphql.ID("0000003"),
				Runs: []MockRun{
					{
						ID:             graphql.ID("3"),
						StartTimestamp: 200,
						Duration:       100,
						Status:         "IN_PROGRESS",
						Log:            "",
					},
				},
			},
		},
	},
}

func GetRepoNames() []string {
	var res = []string{}
	for _, repo := range MockRepoData {
		res = append(res, repo.Name)
	}
	return res
}

type RepoInfo struct {
	Name      string
	Branches  []BranchInfo
	CommitIDs []graphql.ID
}

type BranchInfo struct {
	Name     string
	CommitID graphql.ID
}

func GetRepoInfo(name string) *RepoInfo {
	for _, repo := range MockRepoData {
		if repo.Name == name {
			res := RepoInfo{
				Name: name,
			}
			for _, b := range repo.Branches {
				res.Branches = append(res.Branches, BranchInfo{Name: b.Name, CommitID: b.CommitID})
			}
			for _, c := range repo.Commits {
				res.CommitIDs = append(res.CommitIDs, c.ID)
			}
			return &res
		}
	}
	return nil
}

func GetCommitInfo(repoName string, commitID graphql.ID) *[]MockRun {
	for _, repo := range MockRepoData {
		if repo.Name == repoName {
			for _, c := range repo.Commits {
				if c.ID == commitID {
					return &c.Runs
				}
			}
		}
	}
	return nil
}
