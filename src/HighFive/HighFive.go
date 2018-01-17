package HighFive

import (
"database/sql"
"log"
)

// HighFive object struct
type HighFive struct {
	DB        *sql.DB
	TX        *sql.Tx
}

// Creates a new HighFive object using a given database service and Google Cloud Vision API service
func NewHighFive(db *sql.DB) (*HighFive, error) {
	return &HighFive{
		DB: db,
	}, nil
}

// Sets up a new transaction for the database
func (five *HighFive) NewTX() error {
	tx, err := five.DB.Begin()
	if err != nil {
		log.Println("Error setting tx")
		return err
	}
	five.TX = tx
	return nil
}

// Closes a current transaction for the database
func (five *HighFive) CloseTX() error {
	err := five.DB.Close()
	if err != nil {
		return err
	}
	return nil
}
