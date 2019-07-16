package controllers

import (
	"bankingApp/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ManageController struct {
	beego.Controller
}

func (manage *ManageController) Home() {
	manage.TplName = "default/home.tpl"

	o := orm.NewOrm()
	o.Using("default")

	token := manage.GetString("token")
	if token != "" {

		tok, err1 := strconv.Atoi(token)

		if err1 == nil {
			beego.Debug("cast error")
		}

		user := models.User{Id: tok}
		err := o.Read(&user)

		if err == orm.ErrNoRows {
			beego.Debug("No result found.")
		} else if err == orm.ErrMissPK {
			beego.Debug("No primary key found.")
		} else {
			beego.Debug("Utente presente nel db e rimosso")
			o.Delete(&user)
			manage.Home()
		}

	}

	username := manage.GetString("username")
	if username == "" {
		beego.Debug("username is empty")
		return
	}

	password := manage.GetString("password")
	if password == "" {
		beego.Debug("password is empty")
		return
	}

	if username == "admin" && password == "admin" {
		manage.View()
	}

	var users []*models.User
	num, err1 := o.QueryTable("users").All(&users)

	if err1 != orm.ErrNoRows && num > 0 {
		manage.Data["records"] = users
	}

	user := models.User{Username: username, Password: password}

	err := o.Read(&user, "username", "password")

	if err == orm.ErrNoRows {
		beego.Debug("No result found.")
	} else if err == orm.ErrMissPK {
		beego.Debug("No primary key found.")
	} else {
		beego.Debug("Utente presente nel db")
		manage.UserHome()
	}
}

func (manage *ManageController) UserHome() {
	manage.TplName = "manage/userHome.tpl"
}

//View : it's a view made to visualize all the accounts
func (manage *ManageController) View() {

	manage.TplName = "manage/view.tpl"
	//o := orm.NewOrm()
	//o.Using("default")
}

//Add : add a user
func (manage *ManageController) Add() {

	manage.TplName = "manage/add.tpl"

	user := models.User{}

	if err := manage.ParseForm(&user); err != nil {
		beego.Error("Couldn't parse the form. Reason: %s", err)
	} else {

		if user.Username == "" {
			return
		}

		o := orm.NewOrm()
		o.Using("default")

		user.Balance = 0

		id, err := o.Insert(&user)

		if err == nil {
			msg := fmt.Sprintf("User inserted with id: %d", id)
			beego.Debug(msg)
			manage.Ctx.Redirect(302, "/home")
		} else {
			msg := fmt.Sprintf("Couldn't insert new user. Reason: %s ", err)
			beego.Debug(msg)
		}
	}
}
