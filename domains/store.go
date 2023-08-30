package domains

import "time"

type Store struct {
	ID        int
	Name      string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
