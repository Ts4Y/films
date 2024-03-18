package repository

import (
	"films/internal/entity/actor"
	"films/internal/entity/film"

	"github.com/jmoiron/sqlx"
)

type Actors interface {
	GetAllActors(tx sqlx.Tx) ([]actor.Actor, error)
	GetActorById(tx sqlx.Tx, actorid int) (actor.Actor, error)
	CreateActor(tx sqlx.Tx, actor actor.Actor) error
	ChangeActorInfo(tx sqlx.Tx, actor actor.Actor) error
	DeleteActorById(tx sqlx.Tx, actorid int) error
	GetFilmByActorsId(tx sqlx.Tx, actorid int)([]film.Film,error)
}

type Films interface {
	CreateFilm(tx sqlx.Tx, film film.Film) error
	ChangeFilmInfo(tx sqlx.Tx, film film.Film) error
	DeleteFilmById(tx sqlx.Tx, filmid int) error
	GetFilmByTitle(tx sqlx.Tx, title string)([]film.Film,error)
	GetFilmByActor(tx sqlx.Tx ,actName string)([]film.Film,error)

}
