package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"learn-fiber/config"
	"learn-fiber/internal/dto"
	"learn-fiber/internal/model"
	"learn-fiber/internal/repository"
	"learn-fiber/internal/task/producer"
	"learn-fiber/pkg/crypto"
	"learn-fiber/pkg/redis"
	"time"

	"gorm.io/gorm"
)

type Auth interface {
	Logout(token string) error
	Info(username string) (model.User, error)
	Login(ip string, req dto.LoginRequest) (model.Token, error)
	RenewToken(ip string, req dto.RenewTokenRequest) (model.Token, error)
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

func (h *auth) Login(ip string, req dto.LoginRequest) (token model.Token, err error) {
	key := "attempt:" + req.Username + ":" + ip
	_ = producer.Send(map[string]interface{}{
		"cmd": "send_email",
		"data": map[string]interface{}{
			"title": "Change account password",
			"body":  "Your account password has been changed successfully",
			"to":    "user.Email",
		},
	})
	user, err := h.repo.GetUser(req.Username)
	if err != nil {
		if err = attempts(key); err != nil {
			return token, err
		}
		return token, errors.New("username or password incorrect")
	}
	pass, err := crypto.Decrypt(req.Password)
	if err != nil {
		if err = attempts(key); err != nil {
			return token, err
		}
		return token, errors.New("username or password incorrect")
	}
	if !crypto.Verify(pass, user.Password) {
		if err = attempts(key); err != nil {
			return token, err
		}
		return token, errors.New("username or password incorrect")
	}
	token.Expired = time.Now().Add(time.Hour).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      token.Expired,
	})
	token.Ip = ip
	token.Username = user.Username
	token.RefreshToken = uuid.New().String()
	token.AccessToken, err = claims.SignedString([]byte(config.Cfg.JWT.Secret))
	if err != nil {
		return token, err
	}
	if exist, _ := h.token.Find(ip, req.Username); exist.Id > 0 {
		token.Id = exist.Id
		return h.token.Renew(token)
	}
	return h.token.Add(token)
}

func attempts(key string) error {
	if count, _ := redis.Incr(key); int(count) > config.Cfg.FailedAttempts {
		return errors.New("too many failed attempts")
	} else if count == 1 {
		redis.Exp(key, time.Minute)
	}
	return nil
}

func (h *auth) RenewToken(ip string, req dto.RenewTokenRequest) (token model.Token, err error) {
	token, err = h.token.Get(req.RefreshToken)
	if err != nil {
		return token, errors.New("refresh token invalid")
	}
	user, err := h.repo.GetUser(token.Username)
	if err != nil {
		return token, errors.New("refresh token invalid")
	}
	token.RefreshToken = uuid.NewString()
	token.Expired = time.Now().Add(time.Hour).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      token.Expired,
	})
	token.Ip = ip
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
		return errors.New("user does not exist or has been removed")
	}
	if !crypto.Verify(pass, user.Password) {
		if err = attempts("attempt:" + username); err != nil {
			return err
		}
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
	_ = producer.Send(map[string]interface{}{
		"cmd": "send_email",
		"data": map[string]interface{}{
			"title": "Change account password",
			"body":  "Your account password has been changed successfully",
			"to":    user.Email,
		},
	})
	return h.repo.UpdatePass(username, hashPass)
}
