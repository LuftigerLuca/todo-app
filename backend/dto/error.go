package dto

type ErrorResponse struct {
	Error string `json:"error" example:"Error message" description:"Error description"`
}
