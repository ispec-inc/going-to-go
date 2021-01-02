package mysql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/{{cookiecutter.organization_name}}/{{cookiecutter.repository_name}}/pkg/config"
)

func Init() (*gorm.DB, func() error, error) {
	conn := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		config.RDS.User, config.RDS.Password,
		config.RDS.Host, config.RDS.Port,
		config.RDS.Database,
	)
	db, err := gorm.Open(config.RDS.Driver, conn)
	if err != nil {
		return nil, nil, err
	}

	db.DB().SetMaxIdleConns(config.RDS.MaxIdle)
	db.DB().SetMaxOpenConns(config.RDS.MaxOpen)
	db.DB().SetConnMaxLifetime(config.RDS.MaxLifetime)
	db.LogMode(true)

	return db, cleanup(db), nil
}

func cleanup(db *gorm.DB) func() error {
	return db.Close
}
