package service

import (
	"todo-app/dto"
	"todo-app/model"
)

type ListService struct{}

func (s *ListService) CreateList(req dto.CreateListRequest, userID uint) (uint, error) {
	l := model.List{
		Title:   req.Title,
		OwnerID: userID,
	}

	if len(req.PermittedUsersIDs) > 0 {
		model.DB.Where("id IN ?", req.PermittedUsersIDs).Find(&l.PermittedUsers)
	}

	if err := model.DB.Create(&l).Error; err != nil {
		return 0, err
	}

	return l.ID, nil
}
