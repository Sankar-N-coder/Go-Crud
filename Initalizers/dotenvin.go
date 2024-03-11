package initalizers

import (
	"log"
	"github.com/joho/godotenv"
)
func Inialize(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("An error occured")
	}
}