package dto

type Response[T any] struct {
	StatusCode string `json:"statusCode"`
	Data       T      `json:"data"`
}
