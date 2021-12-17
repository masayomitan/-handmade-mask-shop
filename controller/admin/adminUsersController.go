package controller

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"handmade_mask_shop/component"
	"handmade_mask_shop/domain"
	"handmade_mask_shop/repository"
	"handmade_mask_shop/service"
	"github.com/gin-contrib/sessions"
	_ "github.com/go-sql-driver/mysql"
)

var AdminUser domain.AdminUser

func AdminRegist(c *gin.Context) {
	fmt.Println()

	// db := database.GormConnect()
	// db.Create(&AdminUser)
}


func AdminEdit(c *gin.Context) {
	id := sessions.Default(c).Get("id").(uint)

	adminUser, err := repository.GetAdminUserByID(id)
	if err != nil {
		fmt.Println("error message:", err)
		return
	}
	c.HTML(http.StatusOK, "admin/adminUsers/edit.html", gin.H{
		"adminUser": adminUser,
	})
}


func AdminUpdate (c *gin.Context) {
	id := sessions.Default(c).Get("id").(uint)
	if &id == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session not found"})
	}

	_, err := repository.GetAdminUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "adminUser not found"})
	}

	username := c.PostForm("username")
	password := c.PostForm("password")
	passwordConfirm := c.PostForm("password_confirm")
	if (password != passwordConfirm) {
		fmt.Println("password is not match")
		return
	}
	hash := component.HashPassword(password)
	file, _ := c.FormFile("image")
	
	var imageID uint
	if file != nil {
		newFileName := service.RenameFile(file.Filename)
		imageDir := "./public/images/" 
		filePath := imageDir + newFileName

		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "Unable to save the file",
			})
			return
		}
		imageID = repository.SaveAdminUserImage(newFileName)
	}

	request := map[string]string{
		"username": username,
		"password": hash,
	}

	_, err = repository.UpdateAdminUser(id, imageID, request)
	if err != nil {
		fmt.Println(err)
	}
	c.Redirect(http.StatusFound, "/admin/admin-users/edit")
}