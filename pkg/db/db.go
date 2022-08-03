package db

import (
	"context"
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type Table struct {
	Database    string `db:"database"`
	Name        string `db:"name"`
	Engine      string `db:"engine"`
	IsTemporary bool   `db:"is_temporary"`
}

type Cluster struct {
	Cluster string `db:"cluster"`
}

func MvList(c context.Context) ([]*Table, error) {
	dst := make([]*Table, 0)
	query := `select database, name, engine, is_temporary from system.tables where engine = 'MaterializedView'`
	err := db.SelectContext(c, &dst, query)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

// TableList 查询clickhouse，获取所有的table信息
func TableList(c context.Context) ([]*Table, error) {
	dst := make([]*Table, 0)
	query := `select database, name, engine, is_temporary from system.tables`
	err := db.SelectContext(c, &dst, query)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func ClusterList(c context.Context) (*Cluster, error) {
	dst := Cluster{}
	query := `select cluster from system.clusters`
	err := db.GetContext(c, &dst, query)
	if err != nil {
		return nil, err
	}
	return &dst, nil
}

func Exec(c context.Context, query string) error {
	_, err := db.ExecContext(c, query)
	if err != nil {
		return err
	}
	return nil
}

func InitDB(dsn string) (err error) {
	//dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("clickhouse", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}
