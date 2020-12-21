package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	sparkAppsFieldNames          = builderx.FieldNames(&SparkApps{})
	sparkAppsRows                = strings.Join(sparkAppsFieldNames, ",")
	// 默认代码sparkAppsRowsExpectAutoSet   = strings.Join(stringx.Remove(sparkAppsFieldNames, "id", "create_time", "update_time"), ",")
	sparkAppsRowsExpectAutoSet   = strings.Join(stringx.Remove(sparkAppsFieldNames, "id","create_timestamp","update_time"), ",") // 这里表示不插入"id","create_timestamp","update_time" 这些字段
	sparkAppsRowsWithPlaceHolder = strings.Join(stringx.Remove(sparkAppsFieldNames, "id", "create_time", "update_time"), "=?,") + "=?"
)

type (
	SparkAppsModel interface {
		Insert(data SparkApps) (sql.Result, error)
		FindOne(id int64) (*SparkApps, error)
		FindOneByTime(create_time string) ([]*SparkApps, error)
		Update(data SparkApps) error
		Delete(id int64) error
	}

	defaultSparkAppsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SparkApps struct {
		Id              int64     `db:"id"`
		AppCounts       int64     `db:"app_counts"` // sparkapp数量
		CreateTime      string    `db:"create_time"`
		CreateTimestamp time.Time `db:"create_timestamp"`
	}
)

func NewSparkAppsModel(conn sqlx.SqlConn) SparkAppsModel {
	return &defaultSparkAppsModel{
		conn:  conn,
		table: "sparkApps",
	}
}

func (m *defaultSparkAppsModel) Insert(data SparkApps) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, sparkAppsRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.AppCounts, data.CreateTime)
	return ret, err
}

func (m *defaultSparkAppsModel) FindOne(id int64) (*SparkApps, error) {
	query := fmt.Sprintf("select %s from %s where id = ? limit 1", sparkAppsRows, m.table)
	var resp SparkApps
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
// (*SparkApps, error)
func (m *defaultSparkAppsModel) FindOneByTime(create_time string) ([]*SparkApps, error) {
	query := fmt.Sprintf("select %s from %s where create_time >(?-10)", sparkAppsRows, m.table)
	
	var resp []*SparkApps
	// var resp SparkApps
	err := m.conn.QueryRows(&resp, query, create_time)
	// err := m.conn.QueryRow(&resp, query, create_time)
	fmt.Println("-------[]SparkApps----------")
	fmt.Println(resp)
	fmt.Println(err)

	switch err {
	case nil:
		return resp, nil
		// return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSparkAppsModel) Update(data SparkApps) error {
	query := fmt.Sprintf("update %s set %s where id = ?", m.table, sparkAppsRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.AppCounts, data.CreateTimestamp, data.Id)
	return err
}

func (m *defaultSparkAppsModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
