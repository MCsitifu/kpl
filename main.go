package main

import (
	"github.com/astaxie/beego"
	//_"Beegoproject"
	_ "hallo/routers"
)

func main() {
	//config:=beego.AppConfig
	//appName:=config.String("appname")
	//fmt.Println("应用名称",appName)
	//port.err := config.Int("httpport")
	//if err != nil{
	//	panic("项目配置文件解析失败，请检查配置文件")
	//}
	//dbDriver:=config.String("db_driverNam")
	//dbUser:=config.String("db_user")
	//dbPassword :=config.String("db_password")
	//dbIp :=config.String("db_ip")
	//dbName :=config.String("db_name")
	//
	//connUrl :=dbUser+":"+dbpassword+"@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	//fmt.Println(port)
	//db,err := sql.Open(dbDriver,connUrl)
	beego.Run()

}

