package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func ReadSwitch(name string) *Switch {
	if name == "" {
		return nil
	}
	sw := Switch{Name: name}
	err := orm.NewOrm().Read(&sw, "Name")
	if err != nil {
		beego.Error("Failed to read switch[Name =", sw.Name, "]")
		return nil
	}
	return &sw
}
func DeleteSwitch(name string) error {
	if name == "" {
		return nil
	}
	sw := Switch{Name: name}
	_, err := orm.NewOrm().Delete(&sw, "Name")
	if err != nil {
		beego.Error("Failed to delete switch[Name =", sw.Name, "]")
		return err
	}
	return nil
}

func ReadOrCreateSwitch(name string) (bool, *Switch, error) {
	o := orm.NewOrm()
	sw := Switch{Name: name}

	// return created, id, err
	if created, _, err := o.ReadOrCreate(&sw, "Name"); err == nil {
		if created {
			beego.Debug("New Insert an Switch[Name =", sw.Name, "]")
		} else {
			beego.Debug("Get an Switch[Name =", sw.Name, "]")
		}
		return created, &sw, nil
	} else {
		beego.Error("Failed to Read/Create a Switch[Name =", sw.Name, "]")
		return created, &sw, err
	}
}

// Update Or Create a Switch data structure
func UpdateOrCreateSwitch(this *Switch) error {
	o := orm.NewOrm()
	sw := Switch{Name: this.Name}
	if err := o.Read(&sw, "Name"); err == nil {
		// return num, err
		if _, err := o.Update(this); err == nil {
			beego.Debug("Success to update switch data[Name =", this.Name, "]")
			return nil
		} else {
			beego.Error("Failed to update switch[Name =", this.Name, "] :", err)
			return err
		}
	} else {
		// return id, err
		if _, err := o.Insert(&this); err == nil {
			beego.Debug("Success to Insert switch data[Name =", this.Name, "]")
			return nil
		} else {
			beego.Error("Failed to insert switch[Name =", this.Name, "] :", err)
			return err
		}
	}
}
