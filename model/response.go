package model

type ErrorResponse struct {
	Message string `json:"message"`
}

type DeleteResponse struct {
	ID int64 `json:"id"`
}
