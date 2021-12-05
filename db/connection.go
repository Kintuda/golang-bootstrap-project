package db

import (
	"context"

	"github.com/Kintuda/golang-bootstrap-project/config"
	"github.com/jackc/pgx/v4"
)

type DatabaseConnection struct {
	conn *pgx.Conn
}

func NewDatabaseConnection(c *config.DatabaseConfig) (*DatabaseConnection, error) {
	conn, err := pgx.Connect(context.Background(), c.Dns)

	if err != nil {
		return nil, err
	}

	return &DatabaseConnection{conn: conn}, nil
}

func (d *DatabaseConnection) CloseConnection() {
	d.conn.Close(context.Background())
}
