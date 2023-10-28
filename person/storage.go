package person

import "context"

type Repository interface {
	Create(ctx context.Context, person Person) (int64, error)
	FindAll(ctx context.Context) ([]Person, error)
	FindOne(ctx context.Context, id int64) (Person, error)
	Update(ctx context.Context, person Person) error
	Delete(ctx context.Context, id int64) error
}
