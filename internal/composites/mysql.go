package composites

import (
	"go-clean-news-api/pkg/client/mysql"
	"gorm.io/gorm"
)

type MySQLComposite struct {
	db *gorm.DB
}

func NewMySQLComposite(connString string) *MySQLComposite{
	client, _ := mysql.NewClient(connString)
	return &MySQLComposite{
		db: client,
	}
}