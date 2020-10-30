package repositories

import (
	"database/sql"
	"miao/datamodels"
)

type IProduct interface {
	Conn () error
	Insert(*datamodels.Product) (int64,error)
	Update(*datamodels.Product) error
	Delete(int64) bool
	SelectByKey(int64) *datamodels.Product
	SelectAll() []*datamodels.Product
}

type ProductManager struct {
	Table string
	MysqlConn  *sql.DB
}

// 实例化productManager
func NewProductManager(table string , db *sql.DB) *ProductManager {
	return &ProductManager{
		Table:table,
		MysqlConn:db,
	}
}

func (p *ProductManager) Conn() error {
	return nil
}
