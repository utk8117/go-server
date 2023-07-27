package controller

import (
	"assignment/src/db"
	"assignment/src/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TrackUsers(ctx *gin.Context) {
	fmt.Println("Track user controller")
	var user models.UserBody
	err := ctx.Bind(&user)
	if err != nil {
		fmt.Println("error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Check your req body",
		})
		return
	}
	products := db.ReadValue(user)

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Status Ok from controller",
		"products": products,
	})
}
func AddProduct(ctx *gin.Context) {
	fmt.Println("Track add controller")
	var user models.UserBody
	err := ctx.Bind(&user)
	if err != nil {
		fmt.Println("error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Check your req body",
		})
		return
	}
	fmt.Println("body", len(user.ProductsViewed))
	db.AddUser(user)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Status Ok from controller",
	})
}

func CheckUser(ctx *gin.Context) {
	fmt.Println("Track check user controller")
	var user models.UserBody
	err := ctx.Bind(&user)
	if err != nil {
		fmt.Println("error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Check your req body",
		})
		return
	}
	check, err := db.CheckUser(user)
	if err != nil {
		fmt.Println("error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error in finding from DB",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":    "Status Ok from controller",
		"registered": check.Registered,
		"admin":      check.IsAdmin,
	})
}

func NewProduct(ctx *gin.Context) {
	fmt.Println("Track check user controller")
	var product models.Product
	err := ctx.Bind(&product)
	if err != nil {
		fmt.Println("error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Check your req body",
		})
		return
	}
	err = db.AddProduct(product)
	if err != nil {
		fmt.Println("error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error from DB",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Status OK",
	})

}

func GetRecommendations(ctx *gin.Context) {
	var product models.Product
	err := ctx.Bind(&product)
	if err != nil {
		fmt.Println("error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Check your req body",
		})
		return
	}
	products := db.GetType(product)
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Status OK",
		"products": products,
	})
}
