package util

import (
	"database/sql"
	"fmt"
	m "jrb/database-copy/model"
	"log"

	// "log"
	// "os"

	_ "github.com/mattn/go-sqlite3" // runtime
)

// CopyTables ...
func CopyTables(d m.DbsData) {

	db, err := sql.Open(d.Source.Database, d.Source.ConnString)
	if err != nil {
		log.Fatal(err)
	}
	d.Source.Db = db
	defer db.Close()

	db2, err := sql.Open(d.Target.Database, d.Target.ConnString)
	if err != nil {
		log.Fatal(err)
	}
	d.Target.Db = db2
	defer d.Target.Db.Close()

	for _, tbl := range d.Tables {
		copyRecords(d, tbl)
	}

	// tx, err := dbSrc.Begin()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer stmt.Close()
	// for i := 0; i < 100; i++ {
	// 	_, err = stmt.Exec(i, fmt.Sprintf("closing", i))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// tx.Commit()

	// _, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// rows, err = db.Query("select id, name from foo")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	var id int
	// 	var name string
	// 	err = rows.Scan(&id, &name)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(id, name)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

func copyRecords(d m.DbsData, tablename string) {
	fmt.Println("Copy records from " + tablename)

	rows, err := d.Source.Db.Query("select * from " + tablename)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
