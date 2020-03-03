package util

import (
	"database/sql"
	"fmt"
	m "jrb/database-copy/model"
	"log"
	"reflect"
	"strings"

	// "log"
	// "os"

	_ "github.com/mattn/go-sqlite3" // runtime
	_ "gopkg.in/goracle.v2"         // runtime
)

// CopyTables ...
func CopyTables(d m.DbsData) {
	openDbs(d.Source)
	openDbs(d.Target)

	for _, tbl := range d.Tables {
		copyRecords(d, tbl)
	}
}

func openDbs(dbs m.Dbs) {
	connstr := buildConnStr(dbs)
	db, err := sql.Open(m.DbsMap[dbs.Database], connstr)
	if err != nil {
		log.Fatal(err)
	}
	dbs.Db = db
	defer db.Close()
}

func copyRecords(d m.DbsData, tablename string) {

	var sql = "select * from " + tablename
	rows, _ := d.Source.Db.Query(sql)
	columns, _ := rows.Columns()
	colNum := len(columns)

	var values = make([]interface{}, colNum)
	for i := range values {
		var ii interface{}
		values[i] = &ii
	}

	columnNames := make([]string, 0)
	for _, colName := range columns {
		columnNames = append(columnNames, colName)
	}
	allColumnNames := strings.Join(columnNames, ",")

	inserts := make([]string, 0)

	for rows.Next() {
		vals := scanRow(rows, values, columns)
		inserts = append(inserts, fmt.Sprintf("insert info %v (%v) values (%v) ", tablename, allColumnNames, vals))
	}

	for _, str := range inserts {
		fmt.Println(str)
	}
}

func scanRow(rows *sql.Rows, values []interface{}, columns []string) string {
	err := rows.Scan(values...)

	if err != nil {
		fmt.Println(err)
	}

	retvalues := make([]string, 0)

	for i, colName := range columns {
		var rawvalue = *(values[i].(*interface{}))
		var rawtype = reflect.TypeOf(rawvalue)
		fmt.Println(colName, rawtype, rawvalue)

		retvalues = append(retvalues, toStrValue(fmt.Sprintf("%v", rawvalue), fmt.Sprintf("%v", rawtype)))
	}

	return strings.Join(retvalues, ",")
}

func toStrValue(rawStr string, typ string) string {
	if typ == "string" {
		return "'" + rawStr + "'"
	} else {
		return rawStr
	}
}

func buildConnStr(dbs m.Dbs) string {
	if dbs.Database == m.ORA {
		return dbs.Username + "/" + dbs.Password + "@" + dbs.ConnString
	}
	return dbs.ConnString
}
