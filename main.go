package main

import (
	_ "cat-breeds/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

