package middlewares

import (
	"github.com/marcosfrancomarinho/go-first-project/src/domain/interfaces"
	"github.com/marcosfrancomarinho/go-first-project/src/domain/valuesobject"
)

type UserAuthenticatorMiddlewares struct {
	userAuthenticator interfaces.UserAuthenticator
}

func NewUserAuthenticatorMiddlewares(userAuthenticator interfaces.UserAuthenticator) interfaces.HttpControllers {
	return &UserAuthenticatorMiddlewares{userAuthenticator}
}

func (u *UserAuthenticatorMiddlewares) Execute(httpContext interfaces.HttpContext) {
	token, err := valuesobject.NewToken(httpContext.GetToken())
	if err != nil {
		httpContext.SendError(err)
		return
	}

	idUser, err := u.userAuthenticator.ValidateToken(token)
	if err != nil {
		httpContext.SendError(err)
		return
	}

	httpContext.SetIdentifiers("idUser", idUser)
}
