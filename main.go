package main

import (
	"github.com/astaxie/beego"
	"hallo/dv_mysql"

	//_"Beegoproject"
	_ "hallo/routers"
)

func main() {
	dv_mysql.Connect()
	//1.config:=beego.AppConfig
	//appName:=config.String("appname")
	//fmt.Println("应用名称",appName)
	//port.err := config.Int("httpport")
	//if err != nil{
	//	panic("项目配置文件解析失败，请检查配置文件")
	//}
	//2.dbDriver:=config.String("db_driverNam")
	//dbUser:=config.String("db_user")
	//dbPassword :=config.String("db_password")
	//dbIp :=config.String("db_ip")
	//dbName :=config.String("db_name")
	//db
	//连接数据库
	//connUrl :=dbUser+":"+dbpassword+"@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	//fmt.Println(port)
	//db,err := sql.Open(dbDriver,connUrl)
	//if err != nil{//err不为nil,表示连接数据库时出现了错误，程序就在此中断就可以，不用在执行
	//	//早解决，早解决
	//panic：使程序进入一种恐慌状态
    //      panci"数据库连接错误，请检查配置"
	//}
	beego.Run()
	//代码封装：可以将重复的代码或功能相对比较独立的代码，进行封装，以函数形式进行封装，变成
	//一个代码块或是功能包，供使用者惊醒调用。

}

