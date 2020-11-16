package config

import (
	"ci/manager"
	"ci/util"
	"os"
	"path/filepath"
	"time"

	"code.byted.org/golf/ssconf"
	"code.byted.org/gopkg/dbutil/conf"
	"code.byted.org/gopkg/dbutil/gormdb"
	"code.byted.org/gopkg/logs"
	mysql_driver "code.byted.org/gopkg/mysql-driver"
	"code.byted.org/kite/kite"
	"code.byted.org/kv/goredis"
)

func Init() {
	initDB()
	initRedis()
}

func initRedis() {
	opt := goredis.NewOptionWithTimeout(1000*time.Millisecond, 1000*time.Millisecond, 1000*time.Millisecond, 0, 0, 0, 0)
	opt.SetServiceDiscoveryWithConsul()
	var err error
	manager.CodeFactoryRedis, err = goredis.NewClientWithOption(util.RedisCluster, opt)
	if err != nil {
		logs.Error("Fail to init redis client: %s", util.RedisCluster)
		panic(err)
	}
}

func initDB() {
	confFileName := "conf/db.conf"
	confPath := filepath.Join(kite.ConfigDir, confFileName)
	confData, err := ssconf.LoadSsConfFile(confPath)
	if err != nil {
		logs.Info("cannot load conf file, err: %s", err)
	}

	mysql_driver.OpenInterpolation(util.ServiceName)
	mysql_driver.SetPSMCluster(util.ServiceName, os.Getenv("SERVICE_CLUSTER"))

	dbConf := getDBConf(confData, util.DBName, "read")
	manager.CodeFactoryDBRead = gormdb.NewDBHandlerWithOptional(dbConf)

	dbConf = getDBConf(confData, util.DBName, "write")
	manager.CodeFactoryDBWrite = gormdb.NewDBHandlerWithOptional(dbConf)
}

func getDBConf(config map[string]string, dbName string, mode string) *conf.DBOptional {
	prefix := dbName + "_" + mode
	dbConsul := config[prefix+"_consul"]
	dbConf, err := conf.GetDBOptionalByConsul(dbConsul, "", "")
	if err != nil {
		panic(err.Error())
	}
	dbConf.DBCharset = "utf8mb4"
	dbConf.DBName = config[dbName+"_name"]
	dbConf.Timeout = config[prefix+"_timeout"]
	dbConf.ReadTimeout = config[prefix+"_readtimeout"]
	dbConf.WriteTimeout = config[prefix+"_writetimeout"]
	return &dbConf
}
