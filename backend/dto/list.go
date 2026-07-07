package dto

type CreateListRequest struct {
	Title             string `json:"title" binding:"required" example:"Shopping list"`
	PermittedUsersIDs []uint `json:"permitted_users_ids"`
}

type CreateListResponse struct {
	ListID uint `json:"list_id" description:"ID of the newly created list" example:"1"`
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
