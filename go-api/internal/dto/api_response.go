package dto

type ApiResponse[T any] struct {
	Key     string `json:"key"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}
