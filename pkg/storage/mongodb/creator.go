package mongodb

import "time"

//Creator defines the storage of beer in database
type Creator struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
