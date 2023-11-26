package person

import "context"

type Repository interface {
	Create(ctx context.Context, p Person) (int64, error)
	All(ctx context.Context) ([]Person, error)
	One(ctx context.Context, id int64) (Person, error)
	Update(ctx context.Context, p Person) error
	Delete(ctx context.Context, id int64) error
}
