package test

import (
	"errors"
	"github.com/naruebaet/bitkubsdk/mocks"
	"github.com/naruebaet/bitkubsdk/src/model"
	"github.com/naruebaet/bitkubsdk/usecase"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetServerStatus(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		mockResponse := []model.ServerStatusResponse{
			{
				Name:    "Non-secure endpoints",
				Status:  "ok",
				Message: "",
			},
			{
				Name:    "Secure endpoints",
				Status:  "ok",
				Message: "",
			},
		}

		mockRepo := new(mocks.BitkubRepository)
		mockRepo.On("GetServerStatus").Return(mockResponse, nil).Once()
		uc := usecase.NewUsecase(mockRepo)
		resp, err := uc.GetServerStatus()

		assert.NoError(t, err)
		assert.Equal(t, resp, mockResponse)

		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo := new(mocks.BitkubRepository)
		mockRepo.On("GetServerStatus").Return([]model.ServerStatusResponse{}, errors.New("unexpected error nil data")).Once()
		uc := usecase.NewUsecase(mockRepo)
		_, err := uc.GetServerStatus()

		assert.Error(t, err)

		mockRepo.AssertExpectations(t)
	})

}
