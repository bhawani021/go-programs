package main

import (
	"log"
	"sync"
	"time"
)

type dbConnection struct {
	Host      string
	Username  string
	Password  string
}

var (
		dbConn *dbConnection
		once sync.Once
	)

func getInstance(host, username, password string) *dbConnection{

	var lock = &sync.Mutex{}

	if dbConn == nil {
		lock.Lock()
		defer lock.Unlock()

		if dbConn == nil {
			// Create db connection
			dbConn = &dbConnection{host, username, password}
			log.Println("DB connection created")
		} else {
			log.Println("DB connection exists[1]")
		}
	} else {
		log.Println("DB connection exists[2]")
	}

	return dbConn

}

func getInstance1(host, username, password string) *dbConnection {

	if dbConn == nil {
		once.Do(
			func() {
				log.Println("DB connection created")
				dbConn = &dbConnection{host, username, password}
			})
	} else {
		log.Println("DB connection exists")
	}

	return dbConn
}

func main() {
	dbHost, dbUsername, dbPassword := "localhost", "testUser", "testPassword"

	for i := 0; i < 100; i++ {
		go getInstance1(dbHost, dbUsername, dbPassword)
	}

	time.Sleep(10 * time.Second)
}

