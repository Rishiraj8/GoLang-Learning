package controllers

import (
	"go-jwt-auth/models"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var users = map[string]string{}
var secretKey = []byte("your_secret_key")

func SignUp(c *gin.Context) {
	var input models.User
	if err:= c.ShouldBindJSON(&input); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _,exists := users[input.UserName]; exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.PassWord), bcrypt.DefaultCost)
	if err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	users[input.UserName] = string(hashedPassword)
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})

}



func Login(c *gin.Context) {
	var input models.User
	if err:= c.ShouldBindJSON(&input); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, exists := users[input.UserName]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(input.PassWord))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": input.UserName,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})

	// c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}