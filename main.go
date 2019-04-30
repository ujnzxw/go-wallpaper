package main

import (
	_ "go-wallpaper/routers"
	_ "go-wallpaper/wallpaper"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
