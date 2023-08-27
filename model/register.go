package model

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var UserType string

func Register(ctx *gin.Context) error {
	newUser := User{}
	errBind := ctx.ShouldBind(&newUser)
	if errBind != nil {
		log.Println(errBind.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": "Please check the input and try again :(",
		})
		return errBind
	} else {
		UserType = "Registered"

		_, errExec := DB.Exec(`insert into user(user_id,user_firstname, user_lastname, user_email, user_type) SELECT ?,?,?,?,? where not exists(select user_email from user where user_email=?);`,
			newUser.ID,
			newUser.FirstName,
			newUser.LastName,
			newUser.Email,
			UserType,
			newUser.Email)

		if errExec != nil {
			log.Println(errExec.Error())
		}
		ctx.JSON(http.StatusOK, gin.H{
			"SUCCESS": "You are now registered <3",
		})
		return nil
	}

}
