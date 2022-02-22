package controller

import (
	"fmt"
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/sessions/cookie"
	"encoding/json"
	"handmade_mask_shop/component"
	"handmade_mask_shop/domain"
	"handmade_mask_shop/repository"
	"log"
	"os"
	"strings"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)


func LoginTop (c *gin.Context) {
	fmt.Println()
	c.HTML(http.StatusOK, "admin/login/index.html", gin.H{})
}


func Login(c *gin.Context)  {
	session := sessions.Default(c)
	//オプション指定
	session.Options(sessions.Options{
		HttpOnly: true,
		Secure: true, 
		MaxAge: 86400*14,
		Path: "/",
	})
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Validate form input
	if strings.Trim(email, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can not be empty"})
		return
	}

  request := map[string]string{"email" : email, "password" : password}
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


func ResetPassword(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/logins/reset_password.html", gin.H{})
}


func SendEmailResetPassword(c *gin.Context){
	email := c.PostForm("email")
	rand := component.RandString(20)
  url := os.Getenv("URL")
	remindUrl := url + "/admin/set-password" + rand
	fmt.Println(remindUrl)

	var adminUser domain.AdminUser
  adminUser, err := repository.GetAdminUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}
	repository.SetResetKey(adminUser, rand)

	from := mail.NewEmail(os.Getenv("APP_NAME"), os.Getenv("FROM"))
	subject := "Hakutoハンドメイド"
	to := mail.NewEmail(email, email)
	plainTextContent := "mail!"
	htmlContent := "<strong>テストテストテスト</strong>"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err2 := client.Send(message)

	if err2 != nil {
		log.Println(err2)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	c.Redirect(http.StatusFound, "/admin/reset-password-complete")
}


func ResetPasswordComplete(c *gin.Context){
	c.HTML(http.StatusOK, "admin/logins/reset_password_complete.html", gin.H{})
}


func Logout (c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/admin/")
}