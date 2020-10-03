package services

import (
	"graphsProject/cmd/graphs/models"
)

type serviceDAO interface {
	Save([]models.Graph) ([]models.Graph, error)
}

type service struct {
	dao serviceDAO
}

func NewServiceDAO(dao serviceDAO) *service {

	return &service{dao}
}

func (s service) Save(g []models.Graph) ([]models.Graph, error) {

	graph, err := s.dao.Save(g)

	if err != nil {
		return nil, err
	}

	return graph, nil
}
