package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Adapter struct {
	db *sql.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "hexarch"
)

func NewAdapter() *Adapter {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("Error connecting to db %v", err)
	}

	// CONNECTION TEST
	err = db.Ping()

	if err != nil {
		log.Fatalf("Error testing to db %v", err)
	}

	// CREATE TABLE
	query := "CREATE TABLE IF NOT EXISTS arith_history(id VARCHAR(255) PRIMARY KEY,answer INT NOT NULL , operation  VARCHAR(50) NOT NULL,date timestamp NOT NULL)"

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	_, err = db.ExecContext(ctx, query)

	if err != nil {
		log.Fatalf("Error %s when creating table", err)
	}

	log.Printf("Succesfully connected to the database")
	return &Adapter{db: db}
}

func (da Adapter) CloseDbConnection() {
	err := da.db.Close()

	if err != nil {
		log.Fatalf("Error closing to db connection %v", err)
	}
}

func (da Adapter) AddToHistory(answer int32, operation string) error {

	query := "INSERT INTO arith_history(id,answer, operation, date) VALUES ($1,$2,$3,$4)"

	_, err := da.db.Exec(query, uuid.New().String(), answer, operation, time.Now())
	if err != nil {
		log.Printf("Error inserting to database %v", err)
		return err
	}

	return nil

}
