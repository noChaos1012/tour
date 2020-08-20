/**
对 information_schema 数据库中的COLUMNS表进行连接、查询、数据组装等操作
*/
package sql2struct

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//数据库连接核心对象
type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

//数据库连接
type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

//存储COLUMNS表中的信息
type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

var DBTypeToStructType = map[string]string{
	"int":        "int32",
	"tinyint":    "int8",
	"smallint":   "int",
	"mediumint":  "int64",
	"bigint":     "int64",
	"bit":        "int",
	"bool":       "bool",
	"enum":       "string",
	"set":        "string",
	"varchar":    "string",
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"blob":       "string",
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"time":       "time.Time",
	"float":      "float64",
	"double":     "float64",
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

func (m *DBModel) Connect() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=Local",
		m.DBInfo.UserName, m.DBInfo.Password, m.DBInfo.Host, m.DBInfo.Charset)
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	return nil
}

//获取列信息
func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := `
SELECT 
COLUMN_NAME,DATA_TYPE,COLUMN_KEY,IS_NULLABLE,COLUMN_TYPE,COLUMN_COMMENT 
FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?
`
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var columns []*TableColumn
	for rows.Next() {
		var column TableColumn
		err := rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnKey, &column.IsNullable, &column.ColumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}
		fmt.Printf("column received : %v\n", column)
		columns = append(columns, &column)
	}

	return columns, nil
}
