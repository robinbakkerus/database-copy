package main

import (
	"fmt"
	u "jrb/database-copy/util"
)

func main() {
	fmt.Println("Database copy")

	data := u.ReadProps()

	u.CopyTables(data)

}
