package storage

import (
	"context"

	"github.com/Yury132/Golang-Task-1/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage interface {
	GetUsers(ctx context.Context) ([]models.User, error)
}

type storage struct {
	conn *pgxpool.Pool
}

func (s *storage) GetUsers(ctx context.Context) ([]models.User, error) {
	query := "SELECT id, name, email FROM public.service_user"

	rows, err := s.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users = make([]models.User, 0)
	for rows.Next() {
		var user models.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return users, nil
}

func New(conn *pgxpool.Pool) Storage {
	return &storage{
		conn: conn,
	}
}
