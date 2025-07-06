package DTO

type ResponseRegisterUserDTO struct {
	Message string
	IdUser  string
}

func (r *ResponseRegisterUserDTO) ToObject() map[string]any {
	return map[string]any{
		"idUser":  r.IdUser,
		"message": r.Message,
	}
}
