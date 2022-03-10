package middleware

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vabispklp/yap/internal/app/storage/model"
)

const (
	cookieNameUID               = "uid"
	ContextKeyUserID contextKey = "userID"
)

var secret = []byte("Service Secret 1")

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

// AuthHandle middleware авторизации
func AuthHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var value string
		cookie, err := r.Cookie(cookieNameUID)
		if err == nil {
			// если ошибок нет, присваиваем значение иначе пустая строка
			value = cookie.Value
		}

		var (
			id string
			ok bool
		)

		// NOTICE: если неправально понял, то накидывай еще
		if value != "" {
			id, ok, _ = checkUID(value)
			if !ok {
				http.Error(w, errUnauthorized, http.StatusUnauthorized)
				return
			}
		} else {
			id, err = signUp(w)
			if err != nil {
				log.Printf("signUp error: %s", err)
				http.Error(w, errInternal, http.StatusInternalServerError)
				return
			}
		}

		ctx := r.Context()

		next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, ContextKeyUserID, id)))
	})
}

func signUp(w http.ResponseWriter) (string, error) {
	id, err := generateRandomID()
	if err != nil {
		return "", fmt.Errorf("getUIDByUser error: %w", err)
	}

	user := model.User{
		ID: id,
	}
	uid, err := getUIDByUser(user)
	if err != nil {
		return "", fmt.Errorf("getUIDByUser error: %w", err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:  cookieNameUID,
		Value: uid,
	})

	return id, err
}

func getUIDByUser(user model.User) (string, error) {
	sign, err := generateSign([]byte(user.ID))
	if err != nil {
		return "", err
	}

	user.Sign = sign

	uid, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(uid), nil
}

func generateRandomID() (string, error) {
	b := make([]byte, aes.BlockSize)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func checkUID(uid string) (string, bool, error) {
	var user model.User
	data, err := base64.StdEncoding.DecodeString(uid)
	if err != nil {
		return "", false, err
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		return "", false, err
	}

	sign, err := generateSign([]byte(user.ID))
	if err != nil {
		return "", false, err
	}

	return user.ID, bytes.Equal(user.Sign, sign), nil
}

func generateSign(value []byte) ([]byte, error) {
	aesblock, err := aes.NewCipher(secret)
	if err != nil {
		return nil, err
	}

	sign := make([]byte, aes.BlockSize) // зашифровываем
	aesblock.Encrypt(sign, value)

	return sign, nil
}
