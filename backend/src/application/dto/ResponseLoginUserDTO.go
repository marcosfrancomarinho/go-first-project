package dto

type ResponseLoginUserDTO struct {
	Name  string `json:"name"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

