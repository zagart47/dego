package person

import "context"

type Repository interface {
	Create(ctx context.Context, person Person) (string, error)
	FindAll(ctx context.Context) ([]Person, error)
	FindOne(ctx context.Context, id string) (Person, error)
	Update(ctx context.Context, person Person) error
	Delete(ctx context.Context, id string) error
}
