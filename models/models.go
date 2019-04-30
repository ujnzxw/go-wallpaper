package models

import "github.com/astaxie/beego/orm"

type Switch struct {
	Id    int    `orm:"column(id);pk"` // primary key
	Name  string `json:"name"`
	Value string `json:"value"`
	State bool   `json:"state"`
}

func init() {
	orm.RegisterModel(new(Switch))
}
