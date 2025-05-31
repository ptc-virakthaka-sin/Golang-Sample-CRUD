package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"learn-fiber/config"
	"learn-fiber/internal/dto"
	"learn-fiber/internal/model"
	"learn-fiber/internal/repository"
	"learn-fiber/pkg/crypto"
	"time"

	"gorm.io/gorm"
)

type Auth interface {
	Logout(token string) error
	Info(username string) (model.User, error)
	Login(req dto.LoginRequest) (model.Token, error)
	RenewToken(req dto.RenewTokenRequest) (model.Token, error)
	ChangePassword(username string, req dto.ChangePassRequest) error
}

type auth struct {
	token repository.Token
	repo  repository.User
}

func NewAuth(db *gorm.DB) Auth {
	return &auth{
		token: repository.NewToken(db),
		repo:  repository.NewUser(db),
	}
}

func (h *auth) Logout(token string) error {
	return h.token.Remove(token)
}

func (h *auth) Info(username string) (model.User, error) {
	return h.repo.GetUser(username)
}

func (h *auth) Login(req dto.LoginRequest) (token model.Token, err error) {
	user, err := h.repo.GetUser(req.Username)
	if err != nil {
		return token, errors.New("username or password incorrect")
	}
	pass, err := crypto.Decrypt(req.Password)
	if err != nil {
		return token, errors.New("username or password incorrect")
	}
	if !crypto.Verify(pass, user.Password) {
		return token, errors.New("username or password incorrect")
	}
	token.Expired = time.Now().Add(time.Hour * 1).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      token.Expired,
	})
	token.Username = user.Username
	token.RefreshToken = uuid.New().String()
	token.AccessToken, err = claims.SignedString([]byte(config.Cfg.JWT.Secret))
	if err != nil {
		return token, err
	}
	return h.token.Add(token)
}

func (h *auth) RenewToken(req dto.RenewTokenRequest) (token model.Token, err error) {
	token, err = h.token.Get(req.RefreshToken)
	if err != nil {
		return token, errors.New("refresh token invalid")
	}
	user, err := h.repo.GetUser(token.Username)
	if err != nil {
		return token, errors.New("refresh token invalid")
	}
	token.RefreshToken = uuid.NewString()
	token.Expired = time.Now().Add(time.Hour * 1).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      token.Expired,
	})
	token.AccessToken, err = claims.SignedString([]byte(config.Cfg.JWT.Secret))
	if err != nil {
		return token, err
	}
	return h.token.Renew(token)
}

func (h *auth) ChangePassword(username string, req dto.ChangePassRequest) error {
	pass, err := crypto.Decrypt(req.Password)
	if err != nil {
		return err
	}
	user, err := h.repo.GetUser(username)
	if err != nil {
		return errors.New("user has been locked or deleted")
	}
	if !crypto.Verify(pass, user.Password) {
		return errors.New("current password incorrect")
	}
	newPass, err := crypto.Decrypt(req.NewPassword)
	if err != nil {
		return err
	}
	hashPass, err := crypto.Hash(newPass)
	if err != nil {
		return err
	}
	return h.repo.UpdatePass(username, hashPass)
}
