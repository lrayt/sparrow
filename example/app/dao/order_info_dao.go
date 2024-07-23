package dao

import (
	"github.com/lrayt/sparrow/example/Internal/database"
	"github.com/lrayt/sparrow/example/app/model"
)

type OrderInfoDao struct {
	dbm *database.DBManager
}

func NewOrderInfoDao(dbm *database.DBManager) *OrderInfoDao {
	return &OrderInfoDao{
		dbm: dbm,
	}
}

func (d OrderInfoDao) SaveOrderInfo(orderInfo *model.OrderInfo) error {
	return d.dbm.GormDB.Debug().Save(orderInfo).Error
}
