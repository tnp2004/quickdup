package databases

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/tnp2004/quickdup/pkg/databases/databasesException"
)

type Database interface {
	Conn() *sql.Conn
	HealthCheck() error

	ExecTransaction(query string, args []any) error
	QueryRowTransaction(query string, args []any, dest ...any) error
}

func (p *Postgres) ExecTransaction(query string, args []any) error {
	conn := p.Conn()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.Printf("error begin transaction. Error: %s", err.Error())
		return &databasesException.ExecTransaction{}
	}

	if _, err := tx.Exec(query, args...); err != nil {
		log.Printf("error execute. Error: %s", err.Error())
		if err := tx.Rollback(); err != nil {
			log.Printf("error rollback transaction. Error: %s", err.Error())
			return &databasesException.ExecTransaction{}
		}
		log.Println("rollback transaction")
		return &databasesException.ExecTransaction{}
	}

	if err := tx.Commit(); err != nil {
		log.Printf("error commit transaction. Error: %s", err.Error())
		return &databasesException.ExecTransaction{}
	}

	return nil
}

func (p *Postgres) QueryRowTransaction(query string, args []any, dest ...any) error {
	conn := p.Conn()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.Printf("error begin transaction. Error: %s", err.Error())
		return &databasesException.QueryRow{}
	}

	if err := tx.QueryRow(query, args...).Scan(dest...); err != nil {
		log.Printf("error execute. Error: %s", err.Error())
		if err := tx.Rollback(); err != nil {
			log.Printf("error rollback transaction. Error: %s", err.Error())
			return &databasesException.QueryRow{}
		}
		log.Println("rollback transaction")
		return &databasesException.QueryRow{}
	}

	if err := tx.Commit(); err != nil {
		log.Printf("error commit transaction. Error: %s", err.Error())
		return &databasesException.QueryRow{}
	}

	return nil
}
