package main

import (
	_ "github.com/jicg/App/routers"
	"github.com/astaxie/beego"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/net/html"

)

func main() {
	html.EscapeString("sdfasdf")
	beego.Run()
}

