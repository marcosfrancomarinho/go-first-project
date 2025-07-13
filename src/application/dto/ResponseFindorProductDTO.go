package dto

type ResponseFindorProductDTO struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
	Total    float32 `json:"total"`
}
