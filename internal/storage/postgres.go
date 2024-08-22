package storage

import (
    "database/sql"
    "encoding/json"

    _ "github.com/lib/pq"
)

type Database interface {
    SaveEvent(event map[string]interface{}) error
    Close() error
}

type PostgresDB struct {
    db *sql.DB
}

func NewPostgresDB(url string) (*PostgresDB, error) {
    db, err := sql.Open("postgres", url)
    if err != nil {
        return nil, err
    }
    if err := db.Ping(); err != nil {
        return nil, err
    }
    return &PostgresDB{db: db}, nil
}

func (p *PostgresDB) SaveEvent(event map[string]interface{}) error {
    eventJSON, err := json.Marshal(event)
    if err != nil {
        return err
    }

    _, err = p.db.Exec("INSERT INTO webhooks (data) VALUES ($1)", eventJSON)
    return err
}

func (p *PostgresDB) Close() error {
    return p.db.Close()
}