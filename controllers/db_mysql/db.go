package db_mysql

import(
	_"fmt"
	_"github.com/astaxie/beego/session/mysql"
)//地址未补全
//config:=beego.AppConfig
//dbDriver:=config.String("db_driverNam")
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