package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"net/url"
	"zoob-back/internal/auth"
	"zoob-back/internal/types"
)

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

func SignUp(login string, pass string) error {
	passHash, err := auth.EncryptPass(pass)
	if err != nil {
		return err
	}
	queryString := "INSERT INTO todo.public.users(login, password_hash) VALUES ($1, $2)"
	_, err = Database.Exec(context.Background(), queryString, login, passHash)
	if err != nil {
		return err
	}

	return nil
}
func GetPassHash(login string) (string, error) {
	var parsedPassHash string
	queryString := "SELECT password_hash FROM todo.public.users WHERE login = $1"
	err := Database.QueryRow(context.Background(), queryString, login).Scan(&parsedPassHash)
	if err != nil {
		return "", err
	}
	//var pgErr *pgconn.PgError
	//errors.As(err, &pgErr)
	//if pgErr.Code == pgerrcode.UniqueViolation {
	//	//err = fmt.Errorf("duplicate key value violates unique constraint: %w", err)
	//}

	return parsedPassHash, nil
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
	queryString := "SELECT content FROM todo.public.list_item WHERE item_id = $1"
	var selection string
	err := Database.QueryRow(context.Background(), queryString, id).Scan(&selection)
	if err != nil {
		return "", err
	}
	return selection, nil
}
func UpdateListItem(id int, content string) error {
	queryString := "UPDATE todo.public.list_item SET content = $1 WHERE item_id = $2"
	rows, err := Database.Exec(context.Background(), queryString, content, id)
	if err != nil {
		return err
	}
	if rows.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
func DeleteListItem(id int) error {
	queryString := "DELETE FROM todo.public.list_item WHERE item_id = $1"
	rows, err := Database.Exec(context.Background(), queryString, id)
	if err != nil {
		return err
	}
	if rows.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
func GetAll() ([]types.ListItem, error) {
	queryString := "SELECT item_id, content FROM todo.public.list_item ORDER BY item_id"
	rows, err := Database.Query(context.Background(), queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []types.ListItem

	for rows.Next() {
		var item types.ListItem
		if err := rows.Scan(&item.ItemID, &item.Content); err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}
func DeleteAll() error {
	queryString := "TRUNCATE TABLE todo.public.list_item"
	_, err := Database.Exec(context.Background(), queryString)
	if err != nil {
		return err
	}
	return nil
}
