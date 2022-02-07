package middleware

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"

	"github.com/vabispklp/yap/internal/app/model"
)

const (
	cookieNameUID    = "uid"
	contextKeyUserID = "userID"
)

var secret = []byte("Service Secret 1")

func AuthHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var value string
		cookie, err := r.Cookie(cookieNameUID)
		if err == nil {
			value = cookie.Value
		}

		var uid, id string
		var user model.User

		if value != "" {
			if ok, _ := checkUID(value); !ok {
				http.Error(w, errUnauthorized, http.StatusUnauthorized)
				return
			}
		}

		id, err = generateRandomID()
		if err != nil {
			log.Printf("generateRandomID error: %s", err)
			http.Error(w, errInternal, http.StatusInternalServerError)
			return
		}

		user = model.User{
			ID: id,
		}
		uid, err = getUIDByUser(user)
		if err != nil {
			log.Printf("getUIDByUser error: %s", err)
			http.Error(w, errInternal, http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:  cookieNameUID,
			Value: uid,
		})

		ctx := r.Context()
		context.WithValue(ctx, contextKeyUserID, id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
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

func checkUID(uid string) (bool, error) {
	var user model.User
	data, err := base64.StdEncoding.DecodeString(uid)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		return false, err
	}

	sign, err := generateSign([]byte(user.ID))
	if err != nil {
		return false, err
	}

	return bytes.Equal(user.Sign, sign), nil
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
