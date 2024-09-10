package repository

import "context"

type Repository[M any] interface {
	Insert(context context.Context, model *M) error
	GetByID(context context.Context, id int64) (*M, error)
	Close() error
}
