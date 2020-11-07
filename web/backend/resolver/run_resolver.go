package resolver

import graphql "github.com/graph-gophers/graphql-go"

type RunResolver struct {
	id             graphql.ID
	num            int32
	startTimestamp int32
	duration       int32
	status         string
	log            string
}

func (r *RunResolver) Id() graphql.ID {
	return r.id
}

func (r *RunResolver) Num() int32 {
	return r.num
}

func (r *RunResolver) StartTimestamp() int32 {
	return r.startTimestamp
}

func (r *RunResolver) Duration() int32 {
	return r.duration
}

func (r *RunResolver) Status() string {
	return r.status
}

func (r *RunResolver) Log() string {
	return r.log
}
