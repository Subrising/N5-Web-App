package main

import (
"database/sql"
"fmt"
"net/http"
"time"
//"Dex"
//cli "gopkg.in/urfave/cli.v1"
"os"
_ "github.com/lib/pq"
"log"
"runtime"
"path/filepath"
//"api"
)


type DatabaseConfig struct {
	Host       string
	DBName     string
	DBUser     string
	DBPassword string
	NoSSL      bool
}

var (
	dbConfig          DatabaseConfig
)

func main() {
	runFive()
}

func runFive() {
	db := initDB()

	cv, err := initCloudVisionVision()
	if err != nil {
		panic(err)
	}

	dex, err := Dex.NewDex(db, cv)
	if err != nil {
		panic(err)
	}

	api.StartAPI(*dex)

	_, currTestPath, _, _ := runtime.Caller(0)
	testDataPath := filepath.Join(filepath.Dir(currTestPath), "test_data/images/test")
	log.Println("Test Data Path = ", testDataPath)
	//dex.ResizeAndSendFiles(testDataPath)
	//dex.PageDisplayResults()
	dex.PrintOwnerTable()
}

func initDB() *sql.DB {
	dbConfig.DBUser = "user"
	dbConfig.DBPassword ="password"
	dbConfig.Host = "192.168.99.100"
	dbConfig.DBName = "postgres"
	dbConfig.NoSSL = true
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", dbConfig.DBUser, dbConfig.DBPassword, dbConfig.Host, dbConfig.DBName, dbConfig.SSLMode())
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Println("failed to open database", "err", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return db
}

func (d DatabaseConfig) SSLMode() string {
	// Enable by default
	if d.NoSSL == true {
		return "disable"
	}

	return "require"
}



