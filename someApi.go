package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var customValidator *validator.Validate

type someStruct struct {
	Name   string `json:"name" validate:"customValidate"`
	Gender string `json:"gender"`
}

func init() {
	customValidator = validator.New()
	customValidator.RegisterValidation("customValidate", customValidation)
}

func main() {
	router := gin.Default()

	router.GET("/way-nm", normalWayHandler)
	router.GET("/way-v10", customHandler)

	router.Run(":8080")
}

// normal funcs
func normalWayHandler(c *gin.Context) {
	name := c.Query("name")
	gender := c.Query("gender")

	some := someStruct{
		Name:   name,
		Gender: gender,
	}

	_, err := normalValidate(some) // in this Example bool is no need
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, some)
}

func normalValidate(bro someStruct) (bool, error) { // why i return bool? dont know lol
	name := bro.Name
	gender := bro.Gender

	if gender == "girl" {
		switch name {
		case "somsak":
			// log.Println("this ma girl") <-- no need
			return true, nil
		default:
			return false, errors.New("who?")
		}
	}

	return false, errors.New("error gender not found")
}

// v10 funcs
func customHandler(c *gin.Context) {
	name := c.Query("name")
	gender := c.Query("gender")

	some := someStruct{
		Name:   name,
		Gender: gender,
	}

	if err := customValidator.Struct(some); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, some)
}

func customValidation(fl validator.FieldLevel) bool {
	name := fl.Field().String()
	gender := fl.Parent().FieldByName("Gender").Interface().(string)

	if gender == "girl" {
		switch name {
		case "somsak":
			log.Println("this ma girl")
			return true
		default:
			log.Println("who?")
			return false
		}
	}

	log.Println("error gender not found")
	return false
}
