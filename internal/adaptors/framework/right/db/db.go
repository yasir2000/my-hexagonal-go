package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"

	_ "github.com/go-sql-driver/mysql"
)

type Adaptor struct {
	db *sql.DB
}

func NewAdaptor(driverName, dataSourceName string) (*Adaptor, error) {
	// connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}
	// test db connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failure: %v", err)
	}

	return &Adaptor{db: db}, nil
}

func (da Adaptor) CloseDbConnection() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}

func (da Adaptor) AddToHistory(answer int32, operation string) error {
	queryString, args, err := sq.Insert("arith_history").Columns("data", "answer", "operation").
		Values(time.Now(), answer, operation).ToSql()
	if err != nil {
		return err
	}
	_, err = da.db.Exec(queryString, args...)
	if err != nil {
		return err
	}
	// returning nil, means that we dont want to return anything or that we dont have error
	return nil
}
