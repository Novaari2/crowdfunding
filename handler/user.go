package handler

import (
	"crowdfunding/helper"
	"crowdfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	// butuh dependency user service
	userService user.Service
}

// buat new handler menggunakan depemdemcy service
func NewUserHandler(userService user.Service) *userHandler{
	return &userHandler{userService}
}

// buat handlernya
func (h *userHandler) RegisterUser(c *gin.Context){
	// tangkap input dari user
	// mapping input dari user ke struct RegisterUserInput
	// struct di atas kita passing sebagai parameter service

	// menangkap inputan dari user
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	
	if err != nil{
		// format errornya agar menjadi array 
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Register Account Failed",http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil{
		response := helper.ApiResponse("Rwgister Account Failed",http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokentokenteokenteoketn")

	response := helper.ApiResponse("Account has been created",http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

// login handler
func (h *userHandler) Login(c *gin.Context){
	// user memasukkan input (email dan password)
	//input ditangkap handler
	// mapping dari input user ke input struct
	// input struct di passing ke service
	// di service mencari acount user yg ingin login dengan bantuan repository
	// jika ketemu maka cocokkan password

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil{
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Failed Login, User Not Found",http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil{
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Failed Login, User Not Found",http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokentokerwfeojsdjg")

	response := helper.ApiResponse("Successfully Log in",http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context){
	// ada input email dari user
	// input email di mapping ke struct input
	// struct input di passing ke service
	// service akan memanggil repository - email sudah ada atau belum
	// repository - db

	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil{
		errorMessage := gin.H{"errors": "Server Error"}
		response := helper.ApiResponse("Email Checking Failed",http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailibility(input)

	if err != nil{
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Email Checking Failed",http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable{
		metaMessage = "Email is Available"
	}

	response := helper.ApiResponse(metaMessage,http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}