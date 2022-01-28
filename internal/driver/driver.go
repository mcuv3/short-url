package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DbParams struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func OpenDB(params DbParams) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		params.Host, params.Port, params.User, params.Password, params.DBName))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
