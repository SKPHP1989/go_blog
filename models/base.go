package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsnQuery := "charset=utf8&loc=Asia%2FShanghai"
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?" + dsnQuery
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(User))
}

func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}
