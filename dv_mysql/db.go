package dv_mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/beego/bee/generate"
	"golang.org/x/image/math/fixed"
	"hallo/models"
	"os/user"
)

var Db*sql.DB
//在初始化函数中连接数据库
func init()  {
	fmt.Peintln("链接mysql数据库")
	config := beego.AppConfig
	dbDriver:=config.String("db_driverName")
	dbUser:= config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")

	connUrl := dbUser+":"+dbPassword+"@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	db,err := sql.Open(dbDriver,connUrl)
	if err!=nil{
		panic("数据库连接错误，请检查配置")
	}
	//全局变量赋值
	Db=db
	fmt.Println(db)
}

var Db *sql.DB
//将用户信息保存到数据库

func AddUser(u models.User)(int64,error){
	md5Hash := md5.New()
	md5Hash.Write([]byte(u.Password)//未完
	result, err:=Db.Exec("insert into user (name,birthday,address,password")
	"values(?,?,?,?)",u.User,u.Address,u.Brithday,u.Password
	if err != nil{
		return err
	}
	result.RowsAffected//未完
	)
	//execute,   .exe
}
//将用户信息保存到数据库表当中
//func QueryUser(){
//	Db.QueryRow("select * from ")
//}
//func InserUser(group models.User){
//	//将用户密码进行hash脱敏，使用md5计算密码hash值，并储存hash值
//	hashMd5:=md5.New()
//	hashMd5.Write([]byte(user.Password))
//	bytes := hashMd5.Sum(nil)
//	user.password=hex.EncodeToString(byets)
//	fmt.Println("将要保存的用户名",user.Nick)//漏写
//	result,err := Db.Exec("insert  into user(nickpassword) value (?,?)",user.Nick,user.Password)
//if err!=nil{//3
//	return -1,err
//}
//id,err := result.RowsAffected()
//if err !=nil{
//	return -1,err
//}
//	return id,nil
//}
////查询用户
//func QueryUser(){
//	Db.QueryRow("select * from ")
//}