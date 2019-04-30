package wallpaper

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

// Init function in Main()
func init() {

	// database
	dbFile := beego.AppConfig.String("DataBaseFile")
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	err := orm.RegisterDataBase("default", "sqlite3", dbFile)
	if err != nil {
		beego.Error(err)
	}
	// auto create table
	orm.RunSyncdb("default", false, true)

	runMode := beego.AppConfig.String("runmode")
	if runMode == "dev" {
		orm.Debug = true
	}

}

func Run() {

	src := beego.AppConfig.String("wallpaper::source")
	//	interval, _ := beego.AppConfig.Int("wallpaper::UpdateInterval")

	//for t := range time.NewTicker(time.Duration(interval) * time.Second).C {
	//	beego.Debug("Tick at:", t)

	// Fetch Earth as your wallpaper
	if src == "earth" {
		EarthRun()
		//	}

	}
}
