package dto

type Products struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
}

type ResponseFindorProductDTO struct {
	Total    int        `json:"total"`
	Products []Products `json:"all"`
}
