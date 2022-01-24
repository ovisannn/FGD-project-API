package moderators

import "context"

type ModeratorsUseCase struct {
	moderatorsRepo Repository
}

func NewModeratorsUseCase(moderatorsRepository Repository) UseCase {
	return &ModeratorsUseCase{
		moderatorsRepo: moderatorsRepository,
	}
}

func (UseCase *ModeratorsUseCase) GetByCategoryID(ctx context.Context, idCategory string) ([]Domain, error) {
	result, err := UseCase.moderatorsRepo.GetByCategoryID(ctx, idCategory)
	if err != nil {
		return []Domain{}, err
	}
	// fmt.Println(result)s
	return result, nil
}
