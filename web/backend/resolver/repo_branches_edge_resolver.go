package resolver

type RepoBranchesEdgeResolver struct {
	cursor string
	node   *BranchResolver
}

func (r *RepoBranchesEdgeResolver) Cursor() string {
	return r.cursor
}

func (r *RepoBranchesEdgeResolver) Node() *BranchResolver {
	return r.node
}
