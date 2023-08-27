package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sandwichyum/model"
)

func Start() {
	router := gin.Default()
	router.GET("/sandwichyum.com", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, model.GetMenu())
	})
	router.POST("/sandwichyum.com/order", func(ctx *gin.Context) {
		err := model.OrderPlace(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Ops": "error in bind",
			})
		}
	})
	router.POST("/sandwichyum.com/register", func(ctx *gin.Context) {
		err := model.Register(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Ops": "error in bind",
			})
		}
	})
	//router.POST("/sandwichyum.com/delivery")
	router.Run()
}
