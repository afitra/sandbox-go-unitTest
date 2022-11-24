package service

import (
	"testing"

	"github.com/afitra/sandbox-go-unitTest/entity"
	"github.com/afitra/sandbox-go-unitTest/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}

var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}
func TestCategoryService_GetNotSuccess(t *testing.T) {
	data := entity.Category{
		Id:   "2",
		Name: "budi",
	}
	categoryRepository.Mock.On("FindById", "2").Return(data)
	category, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, category)

	assert.Equal(t, category.Id, category.Id)
	assert.Equal(t, category.Name, category.Name)
}
