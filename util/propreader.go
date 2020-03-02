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

	data.Source.Database = p.GetString("srcDbs", "")
	data.Source.ConnString = p.GetString("srcConnStr", "")
	data.Source.Username = p.GetString("srcUsername", "")
	data.Source.Password = p.GetString("srcPassword", "")

	data.Target.Database = p.GetString("trgDbs", "")
	data.Target.ConnString = p.GetString("trgConnStr", "")
	data.Target.Username = p.GetString("trgUsername", "")
	data.Target.Password = p.GetString("trgPassword", "")

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
