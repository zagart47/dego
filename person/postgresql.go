package person

import (
	"context"
	"dego/pkg/client/postgresql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
)

type repository struct {
	client postgresql.Client
}

func (r *repository) Create(ctx context.Context, person Person) (string, error) {
	personQuery := `
		INSERT INTO public.persons (name, surname, patronymic, age, gender)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING (id)`
	if err := r.client.QueryRow(ctx, personQuery, person.Name, person.Surname, person.Patronymic, person.Age, person.Gender).Scan(&person.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where))
		}
	}
	for _, country := range person.Country {
		countryQuery := `
			INSERT INTO countries (person_id, country_id, probability)
			VALUES ($1, $2, $3)`
		if err := r.client.QueryRow(ctx, countryQuery, person.ID, country.CountryId, country.Probability).Scan(&person.ID); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				fmt.Println(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where))
			}
		}
	}

	return person.ID, nil
}

func (r *repository) FindAll(ctx context.Context) (p []Person, err error) {
	personQuery := `
		SELECT id, name, surname, patronymic, age, gender FROM public.persons
		`
	rows, err := r.client.Query(ctx, personQuery)
	if err != nil {
		return nil, err
	}
	persons := make([]Person, 0)
	countryQuery := `
		SELECT person_id, country_id, probability FROM public.countries WHERE person_id = $1
		`
	for rows.Next() {
		var p Person
		err = rows.Scan(&p.ID, &p.Name, &p.Surname, &p.Patronymic, &p.Age, &p.Gender)
		if err != nil {
			return nil, err
		}
		if err = rows.Err(); err != nil {
			return nil, err
		}
		countryRows, err := r.client.Query(ctx, countryQuery, p.ID)
		if err != nil {
			return nil, err
		}
		for countryRows.Next() {
			var c Country
			err = countryRows.Scan(&c.PersonId, &c.CountryId, &c.Probability)
			if err != nil {
				return nil, err
			}
			p.Country = append(p.Country, c)
		}

		persons = append(persons, p)
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
	countryQuery := `
		SELECT person_id, country_id, probability FROM public.countries WHERE person_id = $1
		`
	countryRows, err := r.client.Query(ctx, countryQuery, id)
	if err != nil {
		return Person{}, err
	}
	for countryRows.Next() {
		var c Country
		err = countryRows.Scan(&c.PersonId, &c.CountryId, &c.Probability)
		if err != nil {
			return Person{}, err
		}
		p.Country = append(p.Country, c)
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
