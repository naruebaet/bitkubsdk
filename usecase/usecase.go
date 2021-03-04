package usecase

import (
	"github.com/naruebaet/bitkubsdk"
	"github.com/naruebaet/bitkubsdk/src/model"
)

type IUsecase interface {
	GetServerStatus() ([]model.ServerStatusResponse, error)
}

type Usecase struct {
	Repo bitkubsdk.BitkubRepository
}

func NewUsecase(rp bitkubsdk.BitkubRepository) IUsecase {
	return &Usecase{
		Repo: rp,
	}
}

func (u *Usecase) GetServerStatus() ([]model.ServerStatusResponse, error) {
	resp, err := u.Repo.GetServerStatus()
	if err != nil {
		return []model.ServerStatusResponse{}, err
	}
	return resp, err
}
