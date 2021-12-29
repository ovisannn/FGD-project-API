package categories

import (
	"context"
	"disspace/helpers/messages"
)

type CategoriresUseCase struct {
	categoriesRepo Repository
}

func NewCategoriesUseCase(categoriesRepository Repository) UseCase {
	return &CategoriresUseCase{
		categoriesRepo: categoriesRepository,
	}
}

func (UseCase *CategoriresUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	result, err := UseCase.categoriesRepo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	// fmt.Println(result)s
	return result, nil
}

func (UseCase *CategoriresUseCase) Create(ctx context.Context, data *Domain) (Domain, error) {
	result, err := UseCase.categoriesRepo.Create(ctx, data)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

func (UseCase *CategoriresUseCase) GetByID(ctx context.Context, id string) (Domain, error) {
	result, err := UseCase.categoriesRepo.GetByID(ctx, id)
	if err != nil {
		return Domain{}, messages.ErrInvalidCategoriesID
	}
	return result, nil
}

func (UseCase *CategoriresUseCase) Delete(ctx context.Context, id string) error {
	err := UseCase.categoriesRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (UseCase *CategoriresUseCase) Update(ctx context.Context, data *Domain, id string) error {
	err := UseCase.categoriesRepo.Update(ctx, data, id)
	if err != nil {
		return err
	}
	return nil
}
