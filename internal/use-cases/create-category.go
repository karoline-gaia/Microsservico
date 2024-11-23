package use_cases

import (
	"github.com/karoline-gaia/go-categories-mvc/internal/entities"
	"github.com/karoline-gaia/go-categories-mvc/internal/repositories"
)

type createCategoryUseCase struct {
	repository *repositories.ICategoryRepository
}

func NewCreateCategoryUseCase(repository *repositories.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{repository}

}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	//TODO: persist.entity.to.db
	// log.Println(category)
	err = u.repository.Save(category)

	if err != nil {
		return err
	}
	return nil
}
