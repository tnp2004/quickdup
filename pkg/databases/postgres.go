package databases

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/tnp2004/quickdup/configs"
)

type Postgres struct {
	*sql.DB
}

type Database interface {
	Conn() *sql.Conn
	HealthCheck() error
}

func NewPostgresDB(cfg *configs.Database) Database {
	connStr := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=%s",
		cfg.Name, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Sslmode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("error connect postgres database. Err: %s", err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatal("ping to postgres database failed")
	}

	return &Postgres{db}
}

func (p *Postgres) Conn() *sql.Conn {
	conn, err := p.DB.Conn(context.Background())
	if err != nil {
		log.Println("get database connection failed")
		return nil
	}
	return conn
}

func (p *Postgres) HealthCheck() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return p.PingContext(ctx)
}
