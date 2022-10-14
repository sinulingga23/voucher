package utility

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToMySQL() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_DB_USER"),
		os.Getenv("MYSQL_DB_PASSWORD"),
		os.Getenv("MYSQL_DB_HOST"),
		os.Getenv("MYSQL_DB_PORT"),
		os.Getenv("MYSQL_DB_NAME")) 
	
	if strings.Compare(os.Getenv("APP_ENV"), "test") == 0 {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			os.Getenv("MYSQL_DB_TEST_USER"),
			os.Getenv("MYSQL_DB_TEST_PASSWORD"),
			os.Getenv("MYSQL_DB_TEST_HOST"),
			os.Getenv("MYSQL_DB_TEST_PORT"),
			os.Getenv("MYSQL_DB_TEST_NAME"))
	}
	
	db, err := sql.Open(os.Getenv("MYSQL_DB_DRIVER"), dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}