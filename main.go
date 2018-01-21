package main

//go:generate sqlboiler postgres

import (
	"database/sql"
	"fmt"
	"log"
	"api"
	_ "github.com/lib/pq"
	"HighFive"
)

// Database struct for connecting
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

// Starts program
func main() {
	runFive()
}

// Initialises database and creates a HighFive object which is the middleware. API is then started.
func runFive() {
	db := initDB()

	five, err := HighFive.NewHighFive(db)
	if err != nil {
		panic(err)
	}
	five.InitData()
	api.StartAPI(*five)
}

// Sets up database
func initDB() *sql.DB {
	dbConfig.DBUser = "user"
	dbConfig.DBPassword ="password"
	dbConfig.Host = "0.0.0.0"
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
	if d.NoSSL == true {
		return "disable"
	}

	return "require"
}



