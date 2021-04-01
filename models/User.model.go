package models

import (
	"errors"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"

	encry "gin_msg/encryption"
)

type User struct {
	mgm.DefaultModel `bson:", inline"`
	Username string
	Email string
	Password string
	PhotoUrl string
	Rooms []string
	BlackListTokens []string
}

func NewUser(username string, email string, password string, photoUrl string) (string, error) {
	if err := mgm.Coll(&User{}).First(bson.M{"username": username}, &User{}); err != nil {
		if err.Error() == "mongo: no documents in result" {
			if hashedPassword, err  := encry.Hashing(password); err != nil {
				return "", err
			} else {
				newUser := &User{
					Username: username,
					Email: email,
					Password: hashedPassword,
					PhotoUrl: photoUrl,
					Rooms: []string{},
				}
				if err := mgm.Coll(newUser).Create(newUser); err != nil {
					return "", err
				}
				return newUser.ID.Hex(), nil
			}
		} else {
			return "", err
		}
	} else {
		return "", errors.New("User exists")
	}
}

func NewUserLogin(username string, password string) (string, error) {
	user := &User{}
	if err := mgm.Coll(&User{}).First(bson.M{"username": username}, user); err != nil {
		return "", err
	}
	if err := encry.CompareHash(user.Password, password); err != nil {
		return "", err
	}
	token, err := encry.CreateToken(user.ID.Hex())
	if err != nil {
		return "", err
	}
	return token, nil
}
