package singleton

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

var db *Database
var lock sync.Mutex

func CreateSingleDatabase() *Database {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("Creating single database instance now")
		time.Sleep(1 * time.Second)
		db = &Database{}
		fmt.Println("Database instance created")
	} else {
		fmt.Println("Database instance already created")
	}
	return db
}

func GetDatabaseInstance() *Database {
	return CreateSingleDatabase()
}
