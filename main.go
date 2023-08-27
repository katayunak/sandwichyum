package main

import (
	_ "github.com/go-sql-driver/mysql"
	"sandwichyum/controller"
	"sandwichyum/model"
)

func init() {
	model.Connection()
	model.GetMenu()
}
func main() {
	controller.Start()
}
