package client

import (
	"github.com/jinzhu/gorm"
)

var (
	DBClient  *gorm.DB
)

func init() {
	DBClient, err := gorm.Open("sqlite3", "code_factory.db")
	if err != nil {
		panic("db connect err")
	}
	defer DBClient.Close()
}
