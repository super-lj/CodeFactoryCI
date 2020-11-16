package manager

import (
	"code.byted.org/gopkg/dbutil/gormdb"
	"code.byted.org/kv/goredis"
)

var (
	CodeFactoryDBRead  *gormdb.DBHandler
	CodeFactoryDBWrite *gormdb.DBHandler
	CodeFactoryRedis   *goredis.Client
)