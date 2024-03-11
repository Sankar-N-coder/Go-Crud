package controller

import (
	initalizers "github.com/Sankar-n-coder/go-crud/Initalizers"
	"github.com/Sankar-n-coder/go-crud/models"
	"github.com/gin-gonic/gin"
)

func Createpost(ctx *gin.Context) {
	var data struct {
		Title string
		Body  string
	}
	ctx.Bind(&data)
	post := models.Post{Title: data.Title, Body: data.Body}
	result := initalizers.DB.Create(&post)
	if result.Error != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200, gin.H{
		"post": post,
	})
}
func ReadPost(ctx *gin.Context){
	var posts []models.Post
	initalizers.DB.Find(&posts)
	ctx.JSON(200,gin.H{
		"post":posts,
	})
}
func Showpost(cxt *gin.Context){
	id:=cxt.Param("id")
	var post models.Post
	initalizers.DB.First(&post,id)
	cxt.JSON(200,gin.H{
		"post":post,
	})

}
func Updatepost(ctx *gin.Context){
	var data struct {
		Title string
		Body  string
	}
	ctx.Bind(&data)
	id:=ctx.Param("id")
	var post models.Post
	initalizers.DB.First(&post,id)
	initalizers.DB.Model(&post).Updates(models.Post{
		Title: data.Title,
		Body: data.Body,
	})
	initalizers.DB.First(&post,id)
	ctx.JSON(200,gin.H{
		"post":post,
	})
}
func DeletePost(ctx *gin.Context){
	id:=ctx.Param("id")
	initalizers.DB.Delete(&models.Post{},id)
	ctx.Status(200)
}
