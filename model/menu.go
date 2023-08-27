package model

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

var (
	sandwichName, sandwichDes, itemName []string
	itemQuan, sandwichPrice             []int
	itemMap                             map[string]int
)

func GetMenu() []Sandwich {
	Menu := []Sandwich{}
	rowMenu, errMenu := DB.Query(`select  * from menu;`)
	if errMenu != nil {
		log.Println(errMenu.Error())
	}
	rowItem, errItem := DB.Query(`select  * from refrigerator;`)
	if errItem != nil {
		log.Println(errItem.Error())
	}
	defer rowMenu.Close()
	defer rowItem.Close()

	for rowMenu.Next() {
		var (
			id, price         int
			name, description string
		)

		errScan := rowMenu.Scan(&id, &name, &price, &description)
		if errScan != nil {
			log.Println(errScan.Error())
		}

		newSandwich := Sandwich{
			ID:          id,
			Name:        name,
			Price:       price,
			Description: description,
		}

		Menu = append(Menu, newSandwich)
		sandwichName = append(sandwichName, name)
		sandwichDes = append(sandwichDes, description)
		sandwichPrice = append(sandwichPrice, price)
	}
	for rowItem.Next() {
		var (
			id, quantity int
			name         string
		)

		errScanItem := rowItem.Scan(&id, &name, &quantity)
		if errScanItem != nil {
			log.Println(errScanItem.Error())
		}
		itemName = append(itemName, name)
		itemQuan = append(itemQuan, quantity)
	}
	return Menu
}

func OrderPlace(ctx *gin.Context) error {
	newOrder := Order{}
	errBind := ctx.ShouldBind(&newOrder)
	if errBind != nil {
		log.Println(errBind.Error())
		return errBind
	}

	for i := 0; i < len(itemName); i++ {
		itemMap = map[string]int{
			itemName[i]: itemQuan[i],
		}
		for i, name := range itemName {
			itemMap[name] = itemQuan[i]
		}
	}

	sandwichAvailability := false
	for i := 0; i < len(sandwichName); i++ {
		if newOrder.SandwichName == sandwichName[i] {
			ingredients := strings.Split(sandwichDes[i], ",")
			for j := 0; j < len(ingredients); j++ {
				for name, quan := range itemMap {
					if ingredients[j] == name {
						quan -= 1
						itemMap[name] = quan
						query := `update refrigerator set item_quantity=? where item_name=?;`
						_, errExec := DB.Exec(query, quan, name)
						if errExec != nil {
							log.Println(errExec.Error())
						}
					}
				}
			}
			sandwichAvailability = true
			ctx.JSON(http.StatusOK, gin.H{
				"Your order has been placed!": "Please go to delivery section...",
			})
			break
		}
	}
	if sandwichAvailability == false {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Sorry ": "We don't have this sandwich please check the menu again :(",
		})
	}

	_, errExec := DB.Exec(`insert into user(user_id,user_firstname, user_lastname, user_email) SELECT ?,?,?,? where not exists(select user_email from user where user_email=?);`,
		newOrder.User.ID,
		newOrder.User.FirstName,
		newOrder.User.LastName,
		newOrder.User.Email,
		newOrder.User.Email,
	)
	if errExec != nil {
		log.Println(errExec.Error())
	}
	return nil
}
