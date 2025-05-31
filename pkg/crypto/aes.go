package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/scrypt"
	"io"
	"learn-fiber/config"
)

const (
	saltSize  = 16
	nonceSize = 12
	keySize   = 32
)

func deriveKey(password, salt []byte) ([]byte, error) {
	return scrypt.Key(password, salt, 1<<15, 8, 1, keySize)
}

func Encrypt(plaintext string) (string, error) {
	salt := make([]byte, saltSize)
	nonce := make([]byte, nonceSize)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", err
	}
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	key, err := deriveKey([]byte(config.Cfg.AES.KEY), salt)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	cipherText := gcm.Seal(nil, nonce, []byte(plaintext), nil)
	data := append(append(salt, nonce...), cipherText...)
	return base64.StdEncoding.EncodeToString(data), nil
}

func Decrypt(encoded string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	if len(data) < saltSize+nonceSize {
		return "", errors.New("invalid encrypted data")
	}
	salt := data[:saltSize]
	nonce := data[saltSize : saltSize+nonceSize]
	cipherText := data[saltSize+nonceSize:]
	key, err := deriveKey([]byte(config.Cfg.AES.KEY), salt)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	plain, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}
	return string(plain), nil
}
