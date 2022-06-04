package mysql

import (
	"context"
	"fmt"
	"time"

	"github.com/tocura/go-jwt-authentication/adapter/repository/mysql/entity"
	"github.com/tocura/go-jwt-authentication/pkg/env"
	"github.com/tocura/go-jwt-authentication/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection(cfg env.Database) (*gorm.DB, error) {
	const format = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True"
	credentials := fmt.Sprintf(
		format,
		cfg.MySQL.Username,
		cfg.MySQL.Password,
		cfg.MySQL.Hostname,
		cfg.MySQL.Port,
		cfg.MySQL.Name,
	)

	db, err := gorm.Open(mysql.Open(credentials))
	if err != nil {
		log.Error(context.TODO(), "error to connect to database", err)
		return nil, err
	}

	sqlDB, _ := db.DB()

	sqlDB.SetConnMaxLifetime(time.Second * 10)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	doMigration(db)

	log.Info(context.TODO(), "connection with database established with success")
	return db, nil
}

func doMigration(db *gorm.DB) {
	if !db.Migrator().HasTable(&entity.User{}) {
		err := db.Migrator().CreateTable(&entity.User{})
		if err != nil {
			panic(err)
		}
	}
}
