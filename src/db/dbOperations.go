package db

import (
	"assignment/src/models"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func SetupDB() *gorm.DB {
	USER := "root"
	PASS := "qwerty1234"
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "products"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	var err error
	db, err = gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func GetDB() *gorm.DB {
	return db
}
func AddProduct(product models.Product) error {
	res := db.Create(&product)
	log.Println("rows", res.RowsAffected)
	return res.Error
}

func AddUser(user models.UserBody) {
	for i := 0; i < len(user.ProductsViewed); i++ {
		user := models.User{
			Name:           user.Name,
			Registered:     user.Registered,
			ProductsViewed: user.ProductsViewed[i],
			IsAdmin:        user.IsAdmin,
		}
		userRes := db.Create(&user)
		log.Println("rows user", userRes.Error)
		if userRes.Error != nil {
			break
		}
	}

}

func ReadValue(user models.UserBody) []models.Product {
	var resultUser []models.User
	var product []models.Product
	var finalProducts []models.Product
	db.Find(&resultUser, "name = ?", user.Name)
	for i := 0; i < len(resultUser); i++ {
		db.Find(&product, "id = ?", resultUser[i].ProductsViewed)
		finalProducts = append(finalProducts, product...)
	}

	log.Println("rows", resultUser, finalProducts)
	return finalProducts
}

func CheckUser(user models.UserBody) (models.User, error) {
	var result models.User

	row := db.First(&result, "name = ?", user.Name)
	log.Println("rows", result)
	return result, row.Error
}

func GetType(product models.Product) []models.Product {
	var result []models.Product
	db.Find(&result, "type = ?", product.Type)
	return result
}
