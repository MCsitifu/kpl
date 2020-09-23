package main

import "github.com/astaxie/beego"

import (
	_ "hallo/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

