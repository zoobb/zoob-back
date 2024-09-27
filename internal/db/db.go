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

var Database *pgx.Conn

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
	Database = connection
	return Database
}

func AddToList(content string) error {
	queryString := "INSERT INTO todo.public.list_item(content) VALUES ($1)"
	_, err := Database.Exec(context.Background(), queryString, content)
	if err != nil {
		return err
	}
	return nil
}
func ReadFromList(id int) (string, error) {
	queryString := "SELECT FROM todo.public.list_item WHERE item_id = $1"
	var selection string
	err := Database.QueryRow(context.Background(), queryString, id).Scan(&selection)
	if err != nil {
		return "", err
	}
	return selection, nil
}
func UpdateListItem(id int, content string) error {
	queryString := "UPDATE todo.public.list_item SET content = $1 WHERE item_id = $2"
	_, err := Database.Exec(context.Background(), queryString, content, id)
	if err != nil {
		return err
	}
	return nil
}
func DeleteListItem(id int) error {
	queryString := "DELETE FROM todo.public.list_item WHERE item_id = $1"
	_, err := Database.Exec(context.Background(), queryString, id)
	if err != nil {
		return err
	}
	return nil
}
