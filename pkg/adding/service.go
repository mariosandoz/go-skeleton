package adding

import "log"

type Service interface {
	AddCreator(c ...Creator) error
}

type Repository interface {
	AddCreator(Creator) error
}

type service struct {
	r Repository
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
