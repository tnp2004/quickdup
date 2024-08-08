package databases

import "database/sql"

type Database interface {
	Conn() *sql.Conn
	HealthCheck() error
}
