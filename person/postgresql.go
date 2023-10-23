package person

import (
	"context"
	"dego/pkg/client/postgresql"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
)

type repository struct {
	client postgresql.Client
}

func (r *repository) Create(ctx context.Context, person Person) (string, error) {
	q := `
		INSERT INTO public.persons (name, surname, patronymic, age, gender)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING (id)`
	if err := r.client.QueryRow(ctx, q, person.Name, person.Surname, person.Patronymic, person.Age, person.Gender).Scan(&person.ID); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			fmt.Println(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where))
		}
		return "", err
	}
	return person.ID, nil
}

func (r *repository) FindAll(ctx context.Context) (p []Person, err error) {
	q := `
		SELECT id, name, surname, patronymic, age, gender FROM public.persons
		`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	persons := make([]Person, 0)
	for rows.Next() {
		var p Person
		err = rows.Scan(&p.ID, &p.Name, &p.Surname, &p.Patronymic, &p.Age, &p.Gender)
		if err != nil {
			return nil, err
		}
		persons = append(persons, p)

		if err = rows.Err(); err != nil {
			return nil, err
		}
	}
	return persons, nil
}

func (r *repository) FindOne(ctx context.Context, id string) (Person, error) {
	q := `
		SELECT id, name, surname, patronymic, age, gender FROM public.persons WHERE id = $1
		`
	var p Person
	err := r.client.QueryRow(ctx, q, id).Scan(&p.ID, &p.Name, &p.Surname, &p.Patronymic, &p.Age, &p.Gender)
	if err != nil {
		return Person{}, err
	}

	return p, nil
}

func (r *repository) Update(ctx context.Context, person Person) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(client postgresql.Client) Repository {
	return &repository{client: client}
}
