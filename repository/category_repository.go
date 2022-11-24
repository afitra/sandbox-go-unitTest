package repository

import "github.com/afitra/sandbox-go-unitTest/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
