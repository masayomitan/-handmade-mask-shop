package controller


import (
	"fmt"
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"handmade_mask_shop/component"
	// "handmade_mask_shop/domain"
	"handmade_mask_shop/repository"
	"strings"
)

func LoginTop (c *gin.Context) {
	fmt.Println()
	c.HTML(http.StatusOK, "front/login/index.html", gin.H{})
}


func Login (c *gin.Context) {
	session := sessions.Default(c)
	
	//オプション指定
	session.Options(sessions.Options{
		HttpOnly: true,
		Secure: true, 
		MaxAge: 86400*7,
		Path: "/",
	})

	username := c.PostForm("username")
	password := c.PostForm("password")

	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can not be empty"})
		return
	}

  request := map[string]string{"username" : username, "password" : password}
	adminUser, err := repository.GetLoginAdminUserByRequest(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	hashedPassword := adminUser.Password
	passwordErr := component.CheckPassword(hashedPassword, password)
  if (passwordErr != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	//interface型なので構造体をjsonに置き換える
	jsonAdminUser, err := json.Marshal(adminUser)
	session.Set("adminUser", string(jsonAdminUser))
	session.Set("id", adminUser.ID)
  err = session.Save();
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.Redirect(http.StatusFound, "/admin/items/")
}


func Logout (c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/admin/")
}