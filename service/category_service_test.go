package service

import (
	"testing"
	"belajar-golang-goruutine/repository"
	"belajar-golang-goruutine/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categorySerice = CategoryRepository{Repository: categoryRepository}

func TestCategoryService_Get(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categorySerice.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, err)
} 

func TestCategoryService_GetFound(t *testing.T) {
	data := entity.Category {
		Id: "2",
		Name: "motor",
	}
	categoryRepository.Mock.On("FindById", "2").Return(data)

	category, err := categorySerice.Get("2")

	assert.Nil(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, data.Id, category.Id)
	assert.Equal(t, data.Name, category.Name)
}