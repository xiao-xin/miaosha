package common

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//  连接Mysql数据
func NewMysqlConn() (*sql.DB,error) {
	db,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/root?charset=utf8")
	if err != nil {
		return nil,err
	}
	db.SetMaxOpenConns(1000)
	err =db.Ping()
	return db,err
}

// 从结果集中获取一条数据
func GetResultRow(rows *sql.Rows) map[string]string {
	// 返回所有的列
	columns , _ := rows.Columns()
	// 一行所有列的值
	vals := make([][]byte,len(columns))
	// 一行填充的数据
	scans := make([]interface{},len(columns))
	// scans引用vales,目的就是把数据填充到val中的[]byte里
	for k,_ := range vals {
		scans[k] = &vals[k]
	}
	result :=make(map[string]string)
	for rows.Next() {
		// 将数据写入到scans中
		rows.Scan(scans...)
		for i ,v := range vals {
			result[columns[i]] = string (v)
		}
	}

	return result
}

// 从结果集中获取多条数据
func GetResultRows(rows *sql.Rows) map[int]map[string]string {
	columns , _ := rows.Columns()
	vals :=make([][]byte,len(columns))
	scans:=make([]interface{},len(columns))

	for k,_ := range vals{
		scans[k] = &vals[k]
	}

	i :=0
	result :=make(map[int]map[string]string,len(columns))
	for rows.Next() {
		rows.Scan(scans...)
		row := make(map[string]string)
		for i,v :=range vals {
			key := columns[i]
			row[key] = string(v)
		}
		result[i] = row
		i++
	}
	return result
}
