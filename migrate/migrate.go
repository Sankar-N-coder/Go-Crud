package main

import (
	initalizers "github.com/Sankar-n-coder/go-crud/Initalizers"
	"github.com/Sankar-n-coder/go-crud/models"
)

func init() {
	initalizers.Inialize()
	initalizers.Databasecon()
}
func main(){
	initalizers.DB.AutoMigrate(&models.Post{})
}