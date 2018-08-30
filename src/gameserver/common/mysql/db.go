package mysql

import (
	"database/sql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"leaf/log"
	"strconv"
	"time"
)

var (
	Context *sql.DB
)

func init() {
	var err error
	Context, err = sql.Open("mysql", "root:root@tcp(192.168.22.203:3306)/gamedb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error("mysql 错误 ：%v", err)
	}
}

type DataRow struct {
	Fields map[string]interface{}
}

func (d *DataRow) GetInt64(fieldName string) int64 {
	v, ok := d.Fields[fieldName]
	if ok {
		return v.(int64)
	} else {
		log.Debug("GetInt64>>>%v>>%v", fieldName, d.Fields)
	}
	return 0
}
func (d *DataRow) GetInt32(fieldName string) int32 {
	v, ok := d.Fields[fieldName]
	if ok {
		if v == nil {
			return 0
		}
		value := int32(0)
		switch v.(type) { //多选语句switch
		case []uint8:
			v1, err := strconv.Atoi(string(v.([]uint8)))
			if err != nil {
				return 0
			}
			value = int32(v1)
		default:
			value = int32(v.(int64))
		}

		return value
	} else {
		log.Debug("GetInt32>>>%v>>%v", fieldName, d.Fields)
	}
	return 0
}
func (d *DataRow) GetInt32Default(fieldName string, val int32) int32 {
	v, ok := d.Fields[fieldName]
	if ok {
		if v == nil {
			return val
		}
		value := int32(0)
		switch v.(type) { //多选语句switch
		case []uint8:
			v1, err := strconv.Atoi(string(v.([]uint8)))
			if err != nil {
				return val
			}
			value = int32(v1)
		default:
			value = int32(v.(int64))
		}

		return value
	} else {
		log.Debug("GetInt32Default>>>%v>>%v", fieldName, d.Fields)
	}
	return val
}
func (d *DataRow) GetTime(fieldName string) time.Time {
	v, ok := d.Fields[fieldName]
	if ok {
		value := v.(time.Time)
		return value
	} else {
		log.Debug("GetTime>>>%v>>%v", fieldName, d.Fields)
	}
	return time.Now()
}
func (d *DataRow) GetString(fieldName string) string {
	v, ok := d.Fields[fieldName]
	if ok {
		if v == nil {
			return ""
		}
		value := string(v.([]uint8))
		return value
	} else {
		log.Debug("GetString>>>%v>>%v", fieldName, d.Fields)
	}
	return ""
}
func (d *DataRow) GetBytes(fieldName string) []byte {
	v, ok := d.Fields[fieldName]
	if ok {
		if v == nil {
			return []byte{}
		}
		value := []byte(v.([]uint8))
		return value
	} else {
		log.Debug("GetBytes>>>%v>>%v", fieldName, d.Fields)
	}
	return []byte{}
}
func (d *DataRow) GetDate(fieldName string) int32 {
	v, ok := d.Fields[fieldName]
	if ok {
		value := v.(time.Time)
		var gVer = value.Format("20060102")
		dd, _ := strconv.Atoi(gVer)
		return int32(dd)
	} else {
		log.Debug("GetDate>>>%v>>%v", fieldName, d.Fields)
	}
	return 20060102
}

func CreateDataRow() *DataRow {
	row := &DataRow{}
	row.Fields = make(map[string]interface{})
	return row
}

func Query(strsql string, args ...interface{}) ([]*DataRow, error) {
	stmt, err := Context.Prepare(strsql)
	if err != nil {
		log.Error("mysql query:%v>>%v", strsql, err)
		return nil, err
	}
	rows, err := stmt.Query(args...)

	//log.Debug("mysql query OpenConnections:%v", Context.Stats().OpenConnections)

	if err != nil {
		log.Error("mysql query:%v>>%v", strsql, err)
		return nil, err
	}
	defer func() {
		stmt.Close()
		rows.Close()
	}()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	for i, v := range columns {
		columns[i] = v
	}
	size := len(columns)
	pts := make([]interface{}, size)
	container := make([]interface{}, size)

	for i := range pts {
		pts[i] = &container[i]
	}
	ret := make([]*DataRow, 0)
	for rows.Next() {
		err = rows.Scan(pts...)
		if err != nil {
			return nil, err
		}
		var r = CreateDataRow()
		for i, column := range columns {
			r.Fields[column] = container[i]
		}
		ret = append(ret, r)
	}
	return ret, nil
}
