package usecase

import (
	"fmt"
	"go-transaction/model"
	"go-transaction/utils/security"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthUsecase interface {
	Login(account model.Auth, ctx *gin.Context) (string, error)
	CheckAccount(account model.Auth) (model.Auth, error)
	Logout(ctx *gin.Context)
}

type authUsecase struct {
	custUsecase CustomerUsecase
	admUsecase  AdminUsecase
}

func (a *authUsecase) Login(account model.Auth, ctx *gin.Context) (string, error) {
	session := sessions.Default(ctx)

	currentSession := session.Get("username")
	fmt.Println(currentSession)
	if currentSession != nil {
		return "", fmt.Errorf("already logged in as '%s'", currentSession)
	}

	auth, err := a.CheckAccount(account)
	if err != nil {
		return "", err
	}
	token := security.CreateAccessToken(auth, account.Role)

	// Set session
	session.Set("username", auth.Username)
	session.Save()

	return token, nil
}

func (a *authUsecase) CheckAccount(account model.Auth) (model.Auth, error) {
	switch account.Role {
	case "user":
		return a.custUsecase.FindUsernameAndPassword(account.Username, account.Password)
	case "admin":
		return a.admUsecase.FindUsernameAndPassword(account.Username, account.Password)
	default:
		return model.Auth{}, fmt.Errorf("role '%s' is doesn't exist", account.Role)
	}
}

func (a *authUsecase) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
}

func NewAuthUsecase(custUsecase CustomerUsecase, admUsecase AdminUsecase) AuthUsecase {
	return &authUsecase{custUsecase, admUsecase}
}
