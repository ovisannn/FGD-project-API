package categories

import "context"

type CategoriresUseCase struct {
	categoriesRepo Repository
}

func (UseCase *CategoriresUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	result, err := UseCase.categoriesRepo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}
