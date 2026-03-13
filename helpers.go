package helpers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/pbkdf2"
)

func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func HashPassword(password string, salt string) (string, error) {
	hash := pbkdf2.Key([]byte(password), []byte(salt), 10000, 32, sha256.New)
	return hex.EncodeToString(hash), nil
}

func VerifyPassword(storedHash string, providedPassword string, salt string) bool {
	newHash, err := HashPassword(providedPassword, salt)
	if err != nil {
		log.Println(err)
		return false
	}
	return newHash == storedHash
}

func IsValidHTTPMethod(method string) bool {
	switch strings.ToUpper(method) {
	case http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete:
		return true
	default:
		return false
	}
}

func GetCurrentTime() time.Time {
	return time.Now().UTC()
}

func HandleError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func ValidateString(input string) error {
	if len(strings.TrimSpace(input)) == 0 {
		return errors.New("input string is empty")
	}
	return nil
}