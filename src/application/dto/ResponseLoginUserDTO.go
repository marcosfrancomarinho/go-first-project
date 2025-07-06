package dto

type ResponseLoginUserDTO struct {
	IdUser  string `json:"idUser"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type object struct {
	IdUser  string `json:"idUser"`
	Message string `json:"message"`
}

func (r *ResponseLoginUserDTO) GetToken() string {
	return r.Token
}
func (r *ResponseLoginUserDTO) ToObject() *object {
	return &object{
		IdUser:  r.IdUser,
		Message: r.Message,
	}
}
