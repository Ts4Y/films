package actor

import (
	"films/internal/entity/film"
	"time"
)

type Actor struct {
	ID          int       `json:"id" db:"id" `
	Name        string    `json:"name" db:"name"`
	Gender      string    `json:"gender" db:"gender"`
	DateOfBirth time.Time `json:"date_of_birth" db:"date_of_birth"`
	FilmsList   []film.Film
}
