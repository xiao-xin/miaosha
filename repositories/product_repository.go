package repositories

import (
	"database/sql"
	"fmt"
	"miao/backend/common"
	"miao/datamodels"
	"strconv"
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
func NewProductManager(table string , db *sql.DB) IProduct {
	return &ProductManager{
		Table:table,
		MysqlConn:db,
	}
}

// 数据库连接
func (p *ProductManager) Conn() error {
	if p.MysqlConn == nil {
		db ,err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		p.MysqlConn=db
	}
	if p.Table == "" {
		p.Table ="product"
	}
	return nil
}

// 新增数据
func (p *ProductManager) Insert(product *datamodels.Product)  (productID int64,err error) {
	// 1. 连接数据库
	if err =p.Conn(); err !=nil {
		return 0,nil
	}

	// 2. 设置SQL
	sql :="INSERT product SET name=?,num=?,image=?,ulr=?"
	stmt,err := p.MysqlConn.Prepare(sql)
	if err != nil {
		return 0,nil
	}

	// 3.执行SQL
	result,err:=stmt.Exec(product.Name,product.Num,product.Image,product.Url)
	if err != nil {
		return 0,err
	}
	// 获取ID
	productID,err = result.LastInsertId()
	return
}

// 删除数据，根据产品ID删除
func (p *ProductManager) Delete(productID int64) bool {
	if err := p.Conn() ; err != nil {
		return false
	}
	sql :="DELETE FROM product WHERE id=?"
	stmt,err:= p.MysqlConn.Prepare(sql)
	if err != nil {
		return false
	}
	_,err=stmt.Exec(productID)
	if err != nil {
		return false
	}
	return true
}

// 更新数据，根据主键ID更新
func (p *ProductManager) Update(product *datamodels.Product) (err error) {
	if err =p.Conn(); err != nil {
		return err
	}

	sql := "UPDATE product SET name=?,num=?,image=?,url=? WHERE id="+ strconv.FormatInt(product.ID,10)
 	stmt,err := p.MysqlConn.Prepare(sql)
	if err != nil {
		return err
	}
	_,err = stmt.Exec(product.Name,product.Num,product.Image,product.Url)
	return err
}

// 根据主键ID查找数据
func (p *ProductManager) SelectByKey(productID int64) (product *datamodels.Product){
	product =&datamodels.Product{}
	var  err error
	if err = p.Conn(); err != nil {
		return 
	}
	sql :="SELECT id,name,num,image,url FROM product WHERE id="+strconv.FormatInt(productID,10)
	rows,err:=p.MysqlConn.Query(sql)
	if err != nil {
		return
	}
	result :=common.GetResultRow(rows)
	common.DataToStructByTagSql(result,product)
	return
}

// 查找所有数据
func (p *ProductManager) SelectAll() (products []*datamodels.Product){
	var err error
	if err = p.Conn(); err != nil {
		fmt.Println(err)
		return
	}
	sql := "SELECT id,name,num,image,url FROM product"
	rows, err := p.MysqlConn.Query(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := common.GetResultRows(rows)
	for  _,v:= range result {
		_p := &datamodels.Product{}
		common.DataToStructByTagSql(v,_p)
		products=append(products,_p)
	}
	return
}