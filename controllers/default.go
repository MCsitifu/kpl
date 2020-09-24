package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"hallo/models"
	"io/ioutil"
)

type MainController struct {
	beego.Controller//匿名字段:
}

//方法的重写
//func (c *MainController) Get() {
//	//获取请求数据
//	userName := c.Ctx.Input.Query("user")
//	password := c.Ctx.Input.Query("psd")
//	//使用固定数据进行数据校验
//	//admin  123456
//	if userName != "admin" || password !="123456"{
//		//代表错误处理
//		c.Ctx.ResponseWriter.Write([]byte("对不起，数据校验错误"))
//		return
//	}
//	//校验正确的情况
//	c.Ctx.ResponseWriter.Write([]byte("对不起，数据校验错误"))
//	c.Data["Website"] = "beego.me"
//	c.Data["Email"] = "astaxie@gmail.com"
//	c.TplName = "index.tpl"
//}
//
//该方法用于处理post请求
//func (c *MainController) Post(){
//	name :=c.Ctx.Request.FormValue("name")
//	age := c.Ctx.Request.FormValue(("age"))
//	sex := c.Ctx.Request.e !="21"&&sex!="male"{
//		c.Ctx.ResponseWriter.Write([]byte("数据校验失败"))
//		return
//	}P
//	c.Ctx.WriteString("数据校验成功")
//}
//该方法用于处理post类型的请求FormValue(("ase"))

//	if name != "admin" && age!="21"{
//func (c*MainController) post(){
//	//解析前端提交的json格式的数据
//	var person models.Person
//	dataBytes, err :=ioutil.ReadAll(c.Ctx.Request.Body)
//	if  err !=nil{
//		c.Ctx.WriteString("数据接受错误，请重试")
//		return
//	}
//	err = json.Unmarshal(dataBytes,&person)
//	if  err !=nil{
//		c.Ctx.WriteString("数据接收错误，请重试")
//		return
//	}
//	fmt.Println("姓名",person.Name)
//	fmt.Println("年龄",person.Age)
//	fmt.Println("性别",person.Sex)
//	c.Ctx.WriteString("数据解析成功")
//
//}
func (c *MainController) Post(){
	//解析前端提交的json格式的数据
	var person models.Person
	dataBytes, err :=ioutil.ReadAll(c.Ctx.Request.Body)
	if  err !=nil{
	c.Ctx.WriteString("数据接收错误，请重试")
	return
	}
	err = json.Unmarshal(dataBytes,&person)
	if  err !=nil{
		c.Ctx.WriteString("数据解析错误，请重试")
		return
	}
	fmt.Println("姓名", person.Name)
	fmt.Println("生日", person.Birthday)
	fmt.Println("地址", person.Address)
	fmt.Println("昵称", person.Nick)
	c.Ctx.WriteString("数据解析成功")
}