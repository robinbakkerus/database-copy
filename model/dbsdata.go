package model

import (
	"database/sql"
)

// Dbs ..
type Dbs struct {
	Database   string
	ConnString string
	Username   string
	Password   string
	Db         *sql.DB
}

// DbsData ...
type DbsData struct {
	Source    Dbs
	Target    Dbs
	Tables    []string
	BatchSize int
	Truncate  bool
}
