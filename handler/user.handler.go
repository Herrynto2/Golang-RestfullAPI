package handler

import (
	"fmt"
	"portservices/model"
	"portservices/repository"
	"portservices/utils"

	"github.com/gin-gonic/gin"
)

// Register ...
func Register(c *gin.Context) {
	var user model.UserRegister
	var err error
	c.BindJSON(&user)
	fmt.Println(user)
	if len(user.Email) == 0 || len(user.Password) == 0 {
		utils.ResponseBadRequest(c, false, "Email and password cannot be empty", gin.H{
			"data": "",
		})
		return
	}
	if len(user.Password) < 5 {
		utils.ResponseBadRequest(c, false, "Password must be at least 5 characters long", gin.H{
			"data": ""})
		return
	}
	if user.Password != user.PasswordConfirm {
		utils.ResponseBadRequest(c, false, "Confirm password do not match", gin.H{
			"data": ""})
	}
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		utils.ResponseServerError(c)
		return
	}
	rest, msg := repository.RegisterUser(&user)
	fmt.Println(rest, msg)
	if msg {
		utils.ResponseSuccess(c, true, "Register success, please verified", gin.H{
			"email": user.Email,
			"code":  rest,
		})
	} else {
		utils.ResponseBadRequest(c, false, "Email has been available", gin.H{
			"data": ""})
	}
}

//Login ...
func Login(c *gin.Context) {
	var loginUser model.UserLogin
	var user model.UserLogin
	var err error
	var token string
	c.ShouldBindJSON(&loginUser)
	if len(loginUser.Email) == 0 || len(loginUser.Password) == 0 {
		utils.ResponseBadRequest(c, false, "Email and password cannot be empty", gin.H{
			"data": ""})
		return
	}
	user, err = repository.LoginUser(loginUser.Email)
	fmt.Println(user.ID)
	if err != nil {
		utils.ResponseServerError(c)
		return
	}
	if utils.CheckPasswordHash(loginUser.Password, user.Password) {
		token, err = utils.GetAuthToken(user.ID)
		if err != nil {
			utils.ResponseServerError(c)
			return
		}
		utils.ResponseSuccess(c, true, "Login success", gin.H{
			"user": gin.H{
				"id_user": user.ID,
				"email":   user.Email,
			},
			"token": token,
		})
		return
	}
	utils.ResponseBadRequest(c, false, "Invalid username or password", gin.H{
		"data": ""})
	return
}

// Logout ...
func Logout(c *gin.Context) {
	var user model.UserLogout
	var err error
	var rest bool
	c.BindJSON(&user)
	if len(user.Email) == 0 {
		utils.ResponseBadRequest(c, false, "Please email cannot be empty", gin.H{
			"data": ""})
	}
	rest, err = repository.LogoutUser(user.Email)
	if !rest || err != nil {
		utils.ResponseServerError(c)
		return
	}
	utils.ResponseSuccess(c, true, "Success to logout", gin.H{
		"email": user.Email,
	})
}

//ChangePassword ...
func ChangePassword(c *gin.Context) {
	// var data model.ConfirmData
	// var user model.User
	// var err error
	// var ok bool
	// c.BindJSON(&data)
	// fmt.Println("data", data)
	// if len(data.Email) == 0 || len(data.Password) == 0 || len(data.UUID) == 0 {
	// 	utils.ResponseBadRequest(c, "Please provide email, password, and uuid")
	// }
	// user, err = repository.GetUserByEmail(data.Email)
	// if err != nil {
	// 	utils.ResponseServerError(c)
	// 	return
	// }
	// if user.UUID != data.UUID {
	// 	utils.ResponseBadRequest(c, "Wrong confirmation code")
	// 	return
	// }
	// ok, err = repository.ConfirmAccount(user.ID)
	// if !ok || err != nil {
	// 	utils.ResponseServerError(c)
	// 	return
	// }
	// utils.ResponseSuccess(c, "Success to confirm account", gin.H{
	// 	"email": user.Email,
	// })
}

//VerifyAccount ...
func VerifyAccount(c *gin.Context) {
	var user model.UserVerify
	c.BindJSON(&user)
	// fmt.Println(user)
	if len(user.VerificationCode) == 0 {
		utils.ResponseBadRequest(c, false, "Verification code cannot be empty", gin.H{
			"data": ""})
		return
	}
	res, err := repository.VerifyAccountUser(user.VerificationCode)
	fmt.Println(res)
	if err != nil {
		utils.ResponseBadRequest(c, false, "Verify code is wrong", gin.H{
			"data": ""})
		return
	}
	utils.ResponseSuccess(c, true, "Verify success", gin.H{
		"code": user.VerificationCode,
	})
	return
}
