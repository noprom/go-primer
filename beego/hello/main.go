package main

import (
	"github.com/astaxie/beego"
	_ "github.com/noprom/go-primer/beego/hello/routers"
)

func main() {
	beego.Run()
}
