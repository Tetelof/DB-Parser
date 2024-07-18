package main

import (
	"log"
	"test/database"
	"test/tables"
)

func main() {
	database.Connect()
	var db_tables []string
	tx := database.DB.Raw("show tables;").Find(&db_tables)
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}
	for _, db_table := range db_tables {
		table := tables.New(db_table)
		tx = database.DB.Raw("describe " + db_table ).Find(&table.Columns)
		if tx.Error != nil {
			log.Fatal(tx.Error)
		}
		table.ChangePath("test/")
		table.Run()
	}
}
