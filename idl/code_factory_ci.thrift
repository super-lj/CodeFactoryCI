namespace py code.factory.ci
namespace go code.factory.ci

include "base.thrift"

struct IsTargetRepoUpdatedRequest {
    1:      optional i64 RepoId,
    255:    base.Base Base
}

struct IsTargetRepoUpdatedResponse {
    1:      optional bool IsUpdated,
    255:    base.BaseResp BaseResp
}

struct FetchTargetRepoLastCommitRequest {
    1:      optional i64 RepoId,
    255:    base.Base Base
}

struct CommitStruct {
    1: optional i64 Id,
    2: optional string  Msg,
    3: optional string Author,
    4: optional i64 LastUpdateTime
}

struct FetchTargetRepoLastCommitResonse {
    1:      optional CommitStruct Commit,
    255:    base.BaseResp BaseResp
}

service CodeFactoryCiService {
    IsTargetRepoUpdatedResponse         IsTargetRepoUpdated         (1: IsTargetRepoUpdatedRequest req),
    FetchTargetRepoLastCommitResonse    FetchTargetRepoLastCommit   (1: FetchTargetRepoLastCommitRequest req)
}
