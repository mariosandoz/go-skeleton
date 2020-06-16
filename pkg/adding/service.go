package adding

import "log"

type Service interface {
	AddCreator(c ...Creator) error
	CheckHealth() string
}

type Repository interface {
	AddCreator(Creator) error
}

type service struct {
	r Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) CheckHealth() string {
	return "Go Skeleton is OK"
}

func (s *service) AddCreator(c ...Creator) error {
	var err error

	//Add logic before inserting creator to repository
	for _, v := range c {
		err = s.r.AddCreator(v)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return err
}
