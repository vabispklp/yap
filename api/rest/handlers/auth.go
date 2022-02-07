package handlers

import (
	"crypto/aes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	cookieNameUID = "uid"

	secret = "Service Secret 1"
)

type User struct {
	ID string `json:"ID"`
}

func auth(w http.ResponseWriter, r *http.Request) (*User, error) {
	cookie, err := r.Cookie(cookieNameUID)
	if err != nil {
		return nil, nil
	}

	var uid, id string
	var user User

	if cookie.Value == "" {
		id, err = generateRandomID()
		if err != nil {
			return nil, err
		}

		user = User{
			ID: id,
		}
		uid, err = getUIDByUser(user)
		if err != nil {
			return nil, nil
		}

		http.SetCookie(w, &http.Cookie{
			Name:  cookieNameUID,
			Value: uid,
		})
	} else {
		if ok, _ := checkUID(cookie.Value); !ok {
			id, err = generateRandomID()
			if err != nil {
				return nil, err
			}

			user = User{
				ID: id,
			}

			uid, err = getUIDByUser(user)
			if err != nil {
				return nil, nil
			}

			http.SetCookie(w, &http.Cookie{
				Name:  cookieNameUID,
				Value: uid,
			})
		}
	}

	return &User{
		ID: uid,
	}, nil
}

func getUIDByUser(user User) (string, error) {
	byteUser, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	aesblock, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}

	dst := make([]byte, aes.BlockSize) // зашифровываем
	aesblock.Encrypt(dst, byteUser)
	fmt.Printf("encrypted: %x\n", dst)

	return fmt.Sprintf("encrypted: %x\n", dst), nil
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
	aesblock, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return false, err
	}
	src := make([]byte, aes.BlockSize) // расшифровываем
	aesblock.Decrypt(src, []byte(uid))
	fmt.Printf("decrypted: %s\n", src)

	ok := json.Valid([]byte(fmt.Sprintf("decrypted: %s\n", src)))
	ok = json.Valid(src)

	return ok, nil
}
