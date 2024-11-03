package repository

import "belajar-golang-goruutine/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}