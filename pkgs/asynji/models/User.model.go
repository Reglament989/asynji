package models

import (
	"errors"
	"fmt"

	"github.com/go-bongo/bongo"
	"go.mongodb.org/mongo-driver/bson"
	mbson "gopkg.in/mgo.v2/bson"

	encry "github.com/Reglament989/asynji/pkgs/asynji/encryption"
)

type Invite struct {
	From     string
	To       string
	When     string
	InviteId string
}

type Update struct {
	Invite Invite
}

type User struct {
	bongo.DocumentBase `bson:",inline"`
	FcmTokens          []string
	Username           string
	Email              string
	Password           string
	PhotoUrl           string
	Rooms              []string
	BlackListTokens    []string
	Updates            []Update
	PublicKeys         []string
}

func (u *User) Save() {
	col := Conn.Collection("Users")
	col.Save(u)
}

func NewUser(username string, email string, password string, photoUrl string) (string, error) {
	col := Conn.Collection("Users")
	err := col.FindOne(bson.M{"username": username}, &User{})
	if dnfError, ok := err.(*bongo.DocumentNotFoundError); ok {
		hashedPassword, _ := encry.Hashing(password)
		user := &User{
			FcmTokens:       []string{},
			Username:        username,
			Password:        hashedPassword,
			Email:           email,
			PhotoUrl:        photoUrl,
			Rooms:           []string{},
			BlackListTokens: []string{},
			Updates:         []Update{},
			PublicKeys:      []string{},
		}
		err := col.Save(user)
		if err != nil {
			return "", err
		}
		return user.GetId().Hex(), nil
	} else {
		fmt.Println("real error " + err.Error())
		return "", dnfError
	}
}

func NewUserLogin(username string, password string) (string, string, error) {
	col := Conn.Collection("Users")
	user := &User{}
	if err := col.FindOne(bson.M{"username": username}, user); err != nil {
		return "", "", err
	}
	if err := encry.CompareHash(user.Password, password); err != nil {
		return "", "", err
	}
	token, refresh, err := encry.CreateTokenPair(user.GetId().Hex())
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
	if yeah, _ := StringInSlice(refreshToken, user.BlackListTokens); yeah {
		return "", "", errors.New("token invalid")
	}
	user.BlackListTokens = append(user.BlackListTokens, refreshToken)
	err1 := Conn.Collection("Users").Save(user)
	if err1 != nil {
		return "", "", err1
	}
	token, refresh, err := encry.CreateTokenPair(user.GetId().Hex())
	if err != nil {
		return "", "", err
	}
	return token, refresh, nil
}

func GetUser(userId string) (*User, error) {
	col := Conn.Collection("Users")
	user := &User{}
	id := mbson.ObjectIdHex(userId)
	if err := col.FindById(id, user); err != nil {
		return &User{}, err
	} else {
		return user, nil
	}
}

func GetCountOfAllUsers() (int, error) {
	return Conn.Collection("Users").Collection().Count()
}

func StringInSlice(a string, list []string) (bool, int) {
	for i, b := range list {
		if b == a {
			return true, i
		}
	}
	return false, 0
}
