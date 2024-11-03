package service

import (
	"belajar-golang-goruutine/repository"
	"belajar-golang-goruutine/entity"
	"errors"
)

type CategoryRepository struct {
	Repository *repository.CategoryRepositoryMock
}

func (service *CategoryRepository) Get(id string) (*entity.Category, error) {
	Category := service.Repository.FindById(id)
	if Category == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return Category, nil
	}
}