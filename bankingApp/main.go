package main

import (
	_ "bankingApp/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"bankingApp/models"

	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default",
		"postgres",
		"user=postgres password=postgres host=127.0.0.1 port=5432 dbname=default sslmode=disable")

	orm.RegisterModel(new(models.User))
	//orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}
