package util

import (
	"fmt"
	"os"
	"strings"

	m "jrb/database-copy/model"

	"github.com/magiconair/properties"
)

// ReadProps read
func ReadProps() m.DbsData {
	p := properties.MustLoadFile(rundir()+"/config.properties", properties.UTF8)
	fmt.Println(p)
	data := m.DbsData{}

	data.Source.Database = strings.ToUpper(p.GetString("source.dbs", ""))
	data.Source.ConnString = p.GetString("source.conn.string", "")
	data.Source.Username = p.GetString("source.username", "")
	data.Source.Password = p.GetString("source.password", "")

	data.Target.Database = strings.ToUpper(p.GetString("target.dbs", ""))
	data.Target.ConnString = p.GetString("target.conn.string", "")
	data.Target.Username = p.GetString("target.username", "")
	data.Target.Password = p.GetString("target.password", "")

	data.BatchSize = p.GetInt("batchSize", 1000)
	data.Truncate = p.GetBool("truncate", true)
	tables := p.GetString("tables", "")
	data.Tables = strings.Split(tables, ",")

	return data
}

func rundir() string {
	dir, _ := os.Getwd()
	return dir
}
