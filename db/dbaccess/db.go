package dbaccess

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Mtbcooler/outrun/config"
)

var db *sql.DB
var DatabaseIsBusy = false

func Set(table, id string, value interface{}) error {
	CheckIfDBSet()
	/*value = Compress(value) // compress the input first
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(key), value)
		if err != nil {
			return err
		}
		return nil
	})
	return err*/

	return nil
}

func Get(table, id string) (interface{}, error) {
	CheckIfDBSet()
	var values interface{}
	err := db.QueryRow("SELECT * FROM ? WHERE id = ?", table, id).Scan(&values)
	if err != nil {
		return nil, err
	}
	return values, nil
}

func Delete(table, id string) error {
	CheckIfDBSet()
	_, err := db.Exec("DELETE FROM ? WHERE id = ?", table, id)
	return err
}

func ForEachKey(bucket string, each func(k, v []byte) error) error {
	CheckIfDBSet()
	// TODO: Reimplement if possible
	return nil
}

func ForEachLogic(each func(tx *sql.Tx) error) error {
	CheckIfDBSet()
	// TODO: Reimplement if possible
	return nil
}

func CheckIfDBSet() {
	if db == nil {
		log.Println("Connecting to MySQL database...")

		sqldb, err := sql.Open("mysql", config.CFile.MySQLUsername+":"+config.CFile.MySQLPassword+"@"+config.CFile.MySQLServerAddress+"/"+config.CFile.MySQLDatabaseName)
		if err != nil {
			log.Println("Failed to open a connection! Check your MySQL settings in config.json for any errors.")
			panic(err)
		}
		err = sqldb.Ping()
		if err != nil {
			log.Println("Failed to connect! Please check your MySQL settings in config.json and try again.")
			panic(err)
		}
		db = sqldb
		log.Println("Successfully connected to database!")
	}
}

func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return errors.New("cannot close database if it's not set!")
}
