package functions

import (
	"database/sql"
	"net/http"
)

func DbConnect(w http.ResponseWriter) (*sql.DB, error) {

	dsn := GetDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {

		return nil, err
	}
	return db, nil

}
