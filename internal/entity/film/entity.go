package film

import "time"

type Film struct {
	ID          int
	Name        string
	Description string
	ReleaseDate time.Time
	Rating float32
}
