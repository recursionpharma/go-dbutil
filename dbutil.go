package dbutil

import (
	"database/sql"
	"fmt"
	"strings"
)

// Connect connects to the database specified by the dbURL.
// It tests the connection by calling db.Ping().
func Connect(dbURL string) (*sql.DB, error) {
	driver, err := GetDriver(dbURL)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open(driver, dbURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// GetDriver extracts the driver from the dbURL;
// for example, it will return postgres for
// postgres://USER:PASSWORD@HOST:PORT/DBNAME
func GetDriver(dbURL string) (string, error) {
	parts := strings.Split(dbURL, "://")
	if len(parts) < 2 {
		return "", fmt.Errorf("DB URL '%s' is missing '://'", dbURL)
	}
	driver := parts[0]
	if driver == "" {
		return "", fmt.Errorf("DB URL '%s' is missing a driver", dbURL)
	}
	return driver, nil
}

// Exists wraps a query with a simple exists check,
//  returning a bool, a la
//    SELECT EXISTS (
//      -- subquery
//    )
//  returning a bool
//
// example usage:
//
//   if Exists("SELECT * FROM foo WHERE bar = 'baz'") {
//		// do something
//   }
func Exists(db *sql.DB, q string, args ...interface{}) (bool, error) {
	var exists bool
	err := db.QueryRow(fmt.Sprintf("SELECT EXISTS(%s)", q), args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return exists, err
	}

	return exists, nil
}
