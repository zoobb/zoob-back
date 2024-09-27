package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"net/url"
)

/*type Database struct {
	User     string
	Password string
	Name     string
	Host     string
}

func (d Database) New() *pgx.Conn {
	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(d.User, d.Password),
		Host:   d.Host,
		Path:   d.Name,
	}

	connection, err := pgx.Connect(context.Background(), u.String())
	if err != nil {
		log.Println("There is an error occurred Connection to the Postgres:", err)
	}
	return connection
}*/

type Credentials struct {
	User     string
	Password string
	Name     string
	Host     string
}

func Connect(credentials Credentials) *pgx.Conn {
	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(credentials.User, credentials.Password),
		Host:   credentials.Host,
		Path:   credentials.Name,
	}
	connection, err := pgx.Connect(context.Background(), u.String())
	if err != nil {
		log.Println("There is an error occurred Connection to the Postgres:", err)
	}
	return connection
}
