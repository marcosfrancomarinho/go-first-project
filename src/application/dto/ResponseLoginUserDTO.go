package dto

type ResponseLoginUserDTO struct {
	IdUser  string `json:"idUser"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

