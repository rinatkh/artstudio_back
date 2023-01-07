package storage

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rinatkh/artstudio_back/config"
)

func InitPsqlDB(c *config.Config) (*sqlx.DB, error) {
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.User,
		c.Postgres.Password,
		c.Postgres.DBName,
		c.Postgres.SSLMode)
	fmt.Println(connectionUrl)
	database, err := sqlx.Connect(c.Postgres.PgDriver, connectionUrl)
	if err != nil {
		return nil, err
	}
	return database, nil
}
