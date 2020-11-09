package mock

type MockRepo struct {
	Name     string
	Branches []MockBranch
	Commits  []MockCommit
}

type MockBranch struct {
	Name     string
	CommitID string
}

type MockCommit struct {
	ID     string
	Msg    string
	Author string
	Runs   []MockRun
}

type MockRun struct {
	Num            int32
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
				CommitID: "0000003",
			},
			{
				Name:     "dev",
				CommitID: "0000002",
			},
		},
		Commits: []MockCommit{
			{
				ID:     "0000001",
				Msg:    "Msg 1",
				Author: "Peixuan Li",
				Runs: []MockRun{
					{
						Num:            1,
						StartTimestamp: 1604701788,
						Duration:       121,
						Status:         "SUCCEED",
						Log:            "",
					},
				},
			},
			{
				ID:     "0000002",
				Msg:    "Msg 2",
				Author: "Peixuan Li",
				Runs: []MockRun{
					{
						Num:            2,
						StartTimestamp: 1604702758,
						Duration:       412,
						Status:         "FAILED",
						Log:            "",
					},
				},
			},
			{
				ID:     "0000003",
				Msg:    "Msg 3",
				Author: "Peixuan Li",
				Runs: []MockRun{
					{
						Num:            3,
						StartTimestamp: 1604703781,
						Duration:       231,
						Status:         "IN_PROGRESS",
						Log: `go run xxxxxx
run 3 succeed!`,
					},
				},
			},
		},
	},
	{
		Name: "test_repo_b",
		Branches: []MockBranch{
			{
				Name:     "main",
				CommitID: "0000006",
			},
			{
				Name:     "dev",
				CommitID: "0000005",
			},
		},
		Commits: []MockCommit{
			{
				ID:     "0000004",
				Msg:    "Msg 4",
				Author: "Xingyou Ji",
				Runs: []MockRun{
					{
						Num:            1,
						StartTimestamp: 1604704281,
						Duration:       321,
						Status:         "SUCCEED",
						Log:            "",
					},
				},
			},
			{
				ID:     "0000005",
				Msg:    "Msg 5",
				Author: "Xingyou Ji",
				Runs: []MockRun{
					{
						Num:            2,
						StartTimestamp: 1604704681,
						Duration:       213,
						Status:         "FAILED",
						Log:            "",
					},
				},
			},
			{
				ID:     "0000006",
				Msg:    "Msg 6",
				Author: "Xingyou Ji",
				Runs: []MockRun{
					{
						Num:            3,
						StartTimestamp: 1604704981,
						Duration:       315,
						Status:         "IN_PROGRESS",
						Log: `go run xxxxxx
run 6 succeed!`,
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
	Name        string
	Branches    []BranchInfo
	CommitHashs []string
}

type BranchInfo struct {
	Name       string
	CommitHash string
}

type CommitInfo struct {
	ID     string
	Msg    string
	Author string
	Runs   []MockRun
}

func GetRepoInfo(name string) *RepoInfo {
	for _, repo := range MockRepoData {
		if repo.Name == name {
			res := RepoInfo{
				Name: name,
			}
			for _, b := range repo.Branches {
				res.Branches = append(res.Branches, BranchInfo{
					Name:       b.Name,
					CommitHash: b.CommitID,
				})
			}
			for _, c := range repo.Commits {
				res.CommitHashs = append(res.CommitHashs, c.ID)
			}
			return &res
		}
	}
	return nil
}

func GetCommitInfo(repoName string, commitID string) *CommitInfo {
	for _, repo := range MockRepoData {
		if repo.Name == repoName {
			for _, c := range repo.Commits {
				if c.ID == commitID {
					return &CommitInfo{
						ID:     c.ID,
						Msg:    c.Msg,
						Author: c.Author,
						Runs:   c.Runs,
					}
				}
			}
		}
	}
	return nil
}
