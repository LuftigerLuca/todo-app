package dto

type CreateListRequest struct {
	Title             string `json:"title" binding:"required"`
	PermittedUsersIDs []uint `json:"permitted_users_ids"`
}

type GetListRequest struct {
	ListID uint `json:"list_id" binding:"required"`
}

type GetListsRequest struct{}

type UpdateListRequest struct {
	ListID            uint   `json:"list_id" binding:"required"`
	Title             string `json:"title"`
	PermittedUsersIDs []uint `json:"permitted_users_ids"`
}

type DeleteListRequest struct {
	ListID uint `json:"list_id" binding:"required"`
}
