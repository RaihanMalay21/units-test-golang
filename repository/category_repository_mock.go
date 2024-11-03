package repository

import (
	"github.com/stretchr/testify/mock"
	"belajar-golang-goruutine/entity"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		Category := arguments.Get(0).(entity.Category)
		return &Category
	}
}