package controllers

import (
	"email-verification/initializers"
	"email-verification/models"
	"email-verification/utils"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm/clause"
	"strings"
)

type RegisterController struct {
	beego.Controller
}

func (register *RegisterController) RegisterUser() {
	register.TplName = "default/register.html"
}

var Email string

func (register *RegisterController) RegisteredUser() {
	var (
		user models.User
	)
	parseFormErr := register.ParseForm(&user)
	utils.CheckError(parseFormErr, "Error parsing form")
	if register.Ctx.Input.Method() == "POST" {
		user.VerifyId = uuid.NewV4()
		user.Name = register.GetString("name")
		user.UserName = register.GetString("username")
		user.Email = register.GetString("email")
		user.Password = register.GetString("password")
		register.SetNames("Name", user.Name)
		register.SetNames("Email", user.Email)
		register.SetNames("Username", user.UserName)
		register.SetNames("Password", user.Password)
		fmt.Println(&user)
		result := initializers.DB.Model(&models.User{}).Clauses(clause.Returning{Columns: []clause.Column{{Name: "user_id"}}}).Create(&user)
		if result.Error != nil {
			if strings.Contains(result.Error.Error(), " duplicate key value violates unique constraint") {
				logs.Info("Duplicate")
				register.TplName = "default/registered.html"
			}
		} else {
			logs.Info("Successfully inserted record", result.Error)
			Email = user.Email
			fmt.Println(Email)
			register.TplName = "default/verify.html"
			utils.SendMail(user.Email, &user)
		}

	}
}
func (register *RegisterController) VerifyRegister() {
	var (
		user models.User
	)
	initializers.DB.Model(&models.User{}).Where("email = ?", Email).Find(&user)
	register.TplName = "default/welcome.html"
	register.SetNames("Id", user.UserId)
	register.SetNames("VerifyId", user.VerifyId)
	register.SetNames("Name", user.Name)
	register.SetNames("Email", Email)
}
func (register *RegisterController) SetNames(name string, value interface{}) {
	register.Data[name] = value
}
