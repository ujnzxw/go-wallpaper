package controllers

import (
	"encoding/json"
	"go-wallpaper/models"
	"go-wallpaper/wallpaper"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
	c.Data["Website"] = "github.com/ujnzxw"
	c.Data["Email"] = "ujnzxw@gmail.com"

	name := "go-wallpaper-switch"

	if created, sw, err := models.ReadOrCreateSwitch(name); err == nil {
		if created {
			beego.Debug("New Created Switch[Switch =", sw, "]")
		} else {
			beego.Debug("Read Existing Switch[Switch =", sw, "]")
		}
		c.Data["GoWallpaperSwitchState"] = sw.State
		beego.Debug("[GoWallpaperSwitchState =", c.Data["GoWallpaperSwitchState"], "]")
	} else {
		beego.Error("Failed to Read Switch[Name =", name, "]")
	}
	wallpaper.Run()
}

func (c *MainController) DoSwitchChange() {
	c.TplName = "index.tpl"
	var sw models.Switch
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &sw)
	if err != nil {
		beego.Error("Failed to decode switch data:[", err, "]")
	} else {
		beego.Debug("Received switch change data: ", sw)
		if err := models.UpdateOrCreateSwitch(&sw); err != nil {
			beego.Error("Failed to save switch data:[", sw, "]")
		}
		s, _ := json.Marshal(sw)
		c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
		c.Ctx.ResponseWriter.Write(s)
	}
}
