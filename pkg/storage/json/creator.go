package json

import "time"

//Creator defines the storage of beer in json file
type Creator struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
