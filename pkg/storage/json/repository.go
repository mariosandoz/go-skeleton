package json

import "go-skeleton/pkg/adding"

// Storage stores creator data in JSON files
type Storage struct {
	creators []Creator
}

//AddCreator saves creator to repository
func (j *Storage) AddCreator(c adding.Creator) error {
	var err error
	nc := Creator{
		Name: c.Name,
	}

	j.creators = append(j.creators, nc)
	return err
}
