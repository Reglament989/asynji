package models

import (
	"errors"
	"log"

	"github.com/kamva/mgm/v3"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"

	encry "asynji/src/encryption"
)

type Update struct {
	Event   string
	Payload interface{}
}

type User struct {
	mgm.DefaultModel `bson:", inline"`
	Id               string
	FcmTokens        []string
	Username         string
	Email            string
	Password         string
	PhotoUrl         string
	Rooms            []string
	BlackListTokens  []string
	Updates          []Update
}

func NewUser(username string, email string, password string, photoUrl string) (string, error) {
	if err := mgm.Coll(&User{}).First(bson.M{"username": username}, &User{}); err != nil {
		if err.Error() == "mongo: no documents in result" {
			if hashedPassword, err := encry.Hashing(password); err != nil {
				return "", err
			} else {
				id := xid.New().String()
				newUser := &User{
					Id:       id,
					Username: username,
					Email:    email,
					Password: hashedPassword,
					PhotoUrl: photoUrl,
					Rooms:    []string{},
				}
				if err := mgm.Coll(newUser).Create(newUser); err != nil {
					return "", err
				}
				return id, nil
			}
		} else {
			return "", err
		}
	} else {
		return "", errors.New("User exists")
	}
}

func NewUserLogin(username string, password string) (string, string, error) {
	user := &User{}
	if err := mgm.Coll(user).First(bson.M{"username": username}, user); err != nil {
		return "", "", err
	}
	if err := encry.CompareHash(user.Password, password); err != nil {
		return "", "", err
	}
	token, refresh, err := encry.CreateTokenPair(user.Id)
	if err != nil {
		return "", "", err
	}
	return token, refresh, nil
}

func RefreshTokens(refreshToken string) (string, string, error) {
	userId, err := encry.VerifyToken(refreshToken)
	if err != nil {
		return "", "", err
	}
	user, err := GetUser(userId)
	if err != nil {
		return "", "", err
	}
	user.BlackListTokens = append(user.BlackListTokens, refreshToken)
	token, refresh, err := encry.CreateTokenPair(user.ID.Hex())
	if err != nil {
		return "", "", err
	}
	return token, refresh, nil
}

func GetUser(userId string) (*User, error) {
	log.Println("Get user", userId)
	user := &User{}
	if err := mgm.Coll(user).First(bson.M{"id": userId}, user); err != nil {
		return &User{}, err
	} else {
		return user, nil
	}
}
