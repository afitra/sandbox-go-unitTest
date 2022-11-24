package service

import (
	"errors"

	"github.com/afitra/sandbox-go-unitTest/entity"
	"github.com/afitra/sandbox-go-unitTest/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {

	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category not found")
	}

	return category, nil

}
