package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/wffranco-demos/go/platzi/rest-api/models"
)

type PostgresUserRepository = PostgresRepository

func (r *PostgresUserRepository) Insert(context context.Context, model *models.User) error {
	_, err := r.db.ExecContext(context, "INSERT INTO users (email, password) VALUES ($1, $2)", model.Email, model.Password)
	return err
}

func (r *PostgresUserRepository) GetByID(context context.Context, id int64) (*models.User, error) {
	row := r.db.QueryRowContext(context, "SELECT id, email, password FROM users WHERE id = $1", id)
	user := new(models.User)
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		log.Fatalf("Error getting user by id: %v\n", err)
	}
	return user, nil
}

func (r *PostgresUserRepository) Close() error {
	return r.db.Close()
}
