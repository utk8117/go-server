package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Publisher   string `json:"publisher"`
	Language    string `json:"language"`
	Edition     int    `json:"edition"`
	Pages       int    `json:"pages"`
	Cost        int    `json:"cost"`
}

type User struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	Name           string `json:"name"`
	Registered     bool   `json:"registered"`
	ProductsViewed int    `json:"products"`
	IsAdmin        bool   `json:"admin"`
}

type UserBody struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	Name           string `json:"name"`
	Registered     bool   `json:"registered"`
	ProductsViewed []int  `json:"products"`
	IsAdmin        bool   `json:"admin"`
}
