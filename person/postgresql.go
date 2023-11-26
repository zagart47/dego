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

func (r *repository) Create(ctx context.Context, p Person) (int64, error) {
	personQuery := `
		INSERT INTO public.persons (name, surname, patronymic, age, gender, is_del)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING (id)`
	if err := r.client.QueryRow(ctx, personQuery, p.Name, p.Surname, p.Patronymic, p.Age, p.Gender, false).Scan(&p.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Printf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
		}
	}
	for _, country := range p.Country {
		countryQuery := `
			INSERT INTO countries (person_id, country_id, probability)
			VALUES ($1, $2, $3)`
		if err := r.client.QueryRow(ctx, countryQuery, p.ID, country.CountryId, country.Probability).Scan(); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				fmt.Printf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
			}
		}
	}

	return p.ID, nil
}

func (r *repository) All(ctx context.Context) ([]Person, error) {
	pq := `
		SELECT id, name, surname, patronymic, age, gender, is_del FROM public.persons		
		`
	rows, err := r.client.Query(ctx, pq)
	if err != nil {
		return nil, err
	}
	var persons = make([]Person, 0)
	for rows.Next() {
		var isDel bool
		var p Person
		if err = rows.Scan(&p.ID, &p.Name, &p.Surname, &p.Patronymic, &p.Age, &p.Gender, &isDel); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				fmt.Printf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
			}
			return nil, err
		}
		if isDel {
			continue
		}
		cq := `
		SELECT country_id, probability FROM public.persons
		JOIN public.countries c on persons.id = c.person_id
		WHERE c.person_id = $1
		`
		countryRows, err := r.client.Query(ctx, cq, p.ID)
		if err != nil {
			return nil, err
		}
		var c Country
		for countryRows.Next() {
			if err = countryRows.Scan(&c.CountryId, &c.Probability); err != nil {
				var pgErr *pgconn.PgError
				if errors.As(err, &pgErr) {
					fmt.Printf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
				}
				return nil, err
			}

			p.Country = append(p.Country, c)
		}
		persons = append(persons, p)
	}
	return persons, nil
}

func (r *repository) One(ctx context.Context, id int64) (Person, error) {
	q := `
		SELECT id, name, surname, patronymic, age, gender, country_id, probability, is_del FROM public.persons
		JOIN public.countries c on persons.id = c.person_id
		WHERE public.persons.id = $1
		`

	rows, err := r.client.Query(ctx, q, id)
	if err != nil {
		return Person{}, err
	}
	var p Person
	for rows.Next() {
		var isDel bool
		var c Country
		if err = rows.Scan(&p.ID, &p.Name, &p.Surname, &p.Patronymic, &p.Age, &p.Gender, &c.CountryId, &c.Probability, &isDel); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				fmt.Printf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
				return Person{}, err
			} else {
				return Person{}, err
			}
		}
		if isDel {
			return Person{}, fmt.Errorf("person not found")
		}
		p.Country = append(p.Country, c)
	}

	return p, nil
}

func (r *repository) Update(ctx context.Context, p Person) error {
	query := `
		UPDATE public.persons SET name = $2, surname = $3, patronymic = $4, age = $5, gender = $6 WHERE id = $1
		`
	if err := r.client.QueryRow(ctx, query, p.ID, p.Name, p.Surname, p.Patronymic, p.Age, p.Gender).Scan(); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Printf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
			return err
		}
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	query := `
		UPDATE public.persons SET is_del = $1 WHERE id = $2
		`
	if err := r.client.QueryRow(ctx, query, true, id).Scan(); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Printf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where)
			return err
		}
	}
	return nil
}

func NewRepository(client postgresql.Client) Repository {
	return &repository{client: client}
}
