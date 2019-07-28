package config

import (
	"fmt"
	"os"
)

func GetDBConnection() string {
	dialect := os.Getenv("DB_DIALECT")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	port := os.Getenv("DB_PORT")
	if dialect == "" || name == "" || pass == "" || user == "" || port == "" || host == "" {
		panic(fmt.Sprintf("One of [DB_DIALECT, DB_NAME, DB_PASS, DB_PORT, DB_USER, DB_HOST] env variables is missing"))
	}

	switch dialect {
	case "mssql":
		return fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, pass, host, port, name)
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)
	}

	return fmt.Sprintf("%s://%s:%s@%s:%s/%s", dialect, user, pass, host, port, name)
}
