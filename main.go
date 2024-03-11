package main

import (
	"fmt"
	"os"
	controller "github.com/Sankar-n-coder/go-crud/Controller"
	initalizers "github.com/Sankar-n-coder/go-crud/Initalizers"
	"github.com/gin-gonic/gin"
)

func init() {
	initalizers.Inialize()
	initalizers.Databasecon()
}
func main() {
	port := os.Getenv("PORT")
	app := gin.Default()
	app.POST("/post", controller.Createpost)
	app.GET("/post",controller.ReadPost)
	app.GET("/post/:id",controller.Showpost)
	app.PUT("/post/update/:id",controller.Updatepost)
	app.DELETE("/post/delete/:id",controller.DeletePost)
	fmt.Println(port)
	app.Run()
}
