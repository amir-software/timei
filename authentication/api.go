package authentication

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	var user User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	userId := user.CreateUser()
	if userId != 0 {
		context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("The user id ==> %d", userId)})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"message": "user not created"})
	}
}

func Login(context *gin.Context) {
	var user User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	is_valid := user.ValidateCredentials()

	if is_valid != true {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}
	token, err := GenerateToken(user.Email, int64(user.ID))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
