package database

import (
	"context"
	"database/sql"
	"github.com/lrayt/sparrow"
	"github.com/lrayt/sparrow/helper"
	"gorm.io/gorm"
	"time"
)

type DBManager struct {
	GormDB *gorm.DB
}

func (p *DBManager) Init() error {
	var (
		err     error
		options = new(helper.DBOptions)
	)
	// cfg
	err = sparrow.GConfigs().PackConf("database.test-db", options)
	if err != nil {
		return err
	}
	// gorm db
	p.GormDB, err = helper.CreateGormDB(options)
	if err != nil {
		return err
	}
	// sql db
	var sqlDB *sql.DB
	sqlDB, err = p.GormDB.DB()
	if err != nil {
		return err
	}
	// ping
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	return sqlDB.PingContext(ctx)
}

func (p DBManager) Close() error {
	db, err := p.GormDB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func NewDBManager() *DBManager {
	return new(DBManager)
}
