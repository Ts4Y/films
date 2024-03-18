package postgresql

import (
	"films/internal/entity/actor"
	"films/internal/entity/film"
	"films/internal/repository"

	"github.com/jmoiron/sqlx"
)

type ActorsRepository struct {
}

func NewActors() repository.Actors {
	return &ActorsRepository{}
}

func (r *ActorsRepository) GetAllActors(tx sqlx.Tx) ([]actor.Actor, error) {
	var actors []actor.Actor

	sqlQuery := `
	select * from actors;
	`

	err := tx.Select(&actors, sqlQuery)
	return actors, err
}

func (r *ActorsRepository) GetActorById(tx sqlx.Tx, actorid int) (actor.Actor, error) {
	actor := actor.Actor{}

	sqlQuery := `
	select * from actors where id = $1;
	`
	err := tx.Get(&actor, sqlQuery, actorid)
	return actor, err
}

func (r *ActorsRepository) CreateActor(tx sqlx.Tx, actor actor.Actor) error {

	sqlQuery := `
	insert into actors (name,gender,date_of_birth),
	values($1,$2,$3)
	`
	_, err := tx.Exec(sqlQuery, actor.Name, actor.Gender, actor.DateOfBirth)

	return err

}

func (r *ActorsRepository) ChangeActorInfo(tx sqlx.Tx, actor actor.Actor) error {

	sqlQuery := `
	update actors set
	name = coalesce(:name,name),
	gender = coalesce(:gender,gender),
	date_of_birth = (:date_of_birth,date_of_birth)
	where id = :id`

	_, err := tx.Exec(sqlQuery, actor.Name, actor.Gender, actor.DateOfBirth, actor.ID)

	return err
}

func (r *ActorsRepository) DeleteActorById(tx sqlx.Tx, actorid int) error {

	sqlQuery := `
	delete from actors
	where id = $1`

	_, err := tx.Exec(sqlQuery, actorid)
	return err
}


func (r *ActorsRepository) GetFilmByActorsId(tx sqlx.Tx, actorid int) ([]film.Film, error) {

	var films []film.Film

	sqlQuery := `
	select actor.name as actor_name, films.name as film_name
	from actors
	join films_list on actors.id = films _list.actor_id
	join films on films_list.film_id = films.id
	order by actors.name; `

	err := tx.Get(films, sqlQuery)

	return films, err
}
