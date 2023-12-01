package adventurer

import (
	model "github.com/arfaghifari/guild-board/src/model/adventurer"
	repo "github.com/arfaghifari/guild-board/src/repository/adventurer"
)

type Usecase interface {
	CreateAdventurer(model.Adventurer) (model.Adventurer, error)
	UpdateAdventurerRank(model.Adventurer) error
	GetAdventurer(int64) (model.Adventurer, error)
}

type usecase struct {
	repo repo.Repository
}

func NewUsecase() (Usecase, error) {
	repo, _ := repo.NewRepository()

	return &usecase{repo}, nil
}

func (u *usecase) CreateAdventurer(adv model.Adventurer) (model.Adventurer, error) {
	return u.repo.CreateAdventurer(adv)
}

func (u *usecase) UpdateAdventurerRank(adv model.Adventurer) error {
	return u.repo.UpdateAdventurerRank(adv)
}

func (u *usecase) GetAdventurer(id int64) (model.Adventurer, error) {
	return u.repo.GetAdventurer(id)
}
