package postgresql

import (
	"films/internal/entity/film"
	"films/internal/repository"

	"github.com/jmoiron/sqlx"
)

type filmsRepository struct {
}

func NewFilms() repository.Films {
	return &filmsRepository{}
}

func (a *filmsRepository) CreateFilm(tx sqlx.Tx, film film.Film) error {
	sqlQuery := `
	insert into films (name, description, release_date, rating),
	values($1,$2,$3,$4)`

	_, err := tx.Exec(sqlQuery, film.Name, film.Description, film.ReleaseDate, film.Rating)

	return err
}

func (a *filmsRepository) ChangeFilmInfo(tx sqlx.Tx, film film.Film) error {
	sqlQuery := `
	update films set
	name = coalesce(:name,name),
	description = coalesce(:description,description),
	release_date = coalesce(:release_date,release_date),
	rating = (:rating,rating)
	where id = :id`
	_, err := tx.Exec(sqlQuery, film.Name, film.Description, film.ReleaseDate, film.Rating, film.ID)

	return err
}

func (a *filmsRepository) DeleteFilmById(tx sqlx.Tx, filmid int) error {

	sqlQuery := `
	delete from films
	where id = $1
	`
	_, err := tx.Exec(sqlQuery, filmid)
	return err
}

func (a *filmsRepository) GetFilmByTitle(tx sqlx.Tx, title string) ([]film.Film, error) {

	var films []film.Film

	sqlQuery := `
	select id,name,description,release_date,rating
	from films
	where lower(name) like '%' || lower($1) || '%'`
	err := tx.Select(&films, sqlQuery)
	return films, err
}

func (a *filmsRepository) GetFilmByActor(tx sqlx.Tx, actName string) ([]film.Film, error) {

	var films []film.Film

	sqlQuery := `
	select f.id,f.name,f.description,f.release_date,f.rating
	from films f
	join films_list fl on (fl.film_id = f.id)
	join actors a on (fl.actor_id = a.id)
	where lower(a.name) like'%' ||lower($1) || '%'`

	err := tx.Select(&films, sqlQuery)
	return films, err
}
