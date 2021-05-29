package routes

import (
	"net/http"

	models "github.com/Reglament989/asynji/pkgs/asynji/models"
	val "github.com/Reglament989/asynji/pkgs/asynji/validators"

	"github.com/gin-gonic/gin"
)

type CreateUserBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// [POST] /user/create
func CreateUserRoute(c *gin.Context) {
	var body CreateUserBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	validation := val.MultiValidation{}
	email := val.Validator{Data: body.Email}
	email.IsEmail()
	validation.Add(email)
	username := val.Validator{Data: body.Username, Name: "username"}
	username.Length(4, 16)
	validation.Add(username)
	password := val.Validator{Data: body.Password, Name: "password"}
	password.Length(8, 32)
	validation.Add(password)
	errors := validation.Result()
	if errors != nil {
		stringsOfErrors := []string{}
		for i := range errors {
			stringsOfErrors = append(stringsOfErrors, errors[i].Error())
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": stringsOfErrors,
		})
		return
	}
	if userId, err := models.NewUser(body.Username, body.Email, body.Password, ""); err != nil {
		c.JSON(409, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"userId": userId,
		})
	}
}

type UploadFcmTokenBody struct {
	Token string `json:"token" binding:"required"`
}

// [PUT] "/user/upload/fcm"
func UploadFcmToken(c *gin.Context) {
	var body UploadFcmTokenBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := models.GetUser(c.GetString("userId"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	user.FcmTokens = append(user.FcmTokens, body.Token)
	user.Save()
	c.JSON(200, gin.H{
		"message": "Saved",
	})
}

type UploadPublicKeyBody struct {
	Key string `json:"key" binding:"required"`
}

// [PUT] "/user/upload/public"
func UploadPublicKey(c *gin.Context) {
	var body UploadPublicKeyBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := models.GetUser(c.GetString("userId"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	user.PublicKeys = append(user.PublicKeys, body.Key)
	user.Save()
	c.JSON(200, gin.H{
		"message": "Saved",
	})
}

// [GET] "/user/:userid/public"
func GetPublicKey(c *gin.Context) {
	userid := c.Param("userid")
	gettenUser, err := models.GetUser(userid)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Successfull",
		"key":     gettenUser.PublicKeys,
	})
}

// [GET] "/user/:userid"
func GetInfoAboutUserRoute(c *gin.Context) {
	userid := c.Param("userid")
	gettenUser, err := models.GetUser(userid)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Successfull",
		"user":    gettenUser,
	})
}
