package pkg_repository

import "go.mongodb.org/mongo-driver/mongo"

type IRepository[T any] interface {
	FindAll() []T
	FindOne(id string) T
	Create(data T) T
	Update(id string, data T) T
	Delete(id string)
}

type SRepository struct {
	Client     *mongo.Client
	Collection string
}

func (s *SRepository) FindAll() ([]interface{}, error) {
	return nil, nil
}
