package models

import (
	"errors"
	"log"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Category struct {
	Id 			int    `orm:"column(id);auto"`  
	Name 		string `orm:"column(title);size(50)"` 
}


func init() {

	orm.RegisterModel(new(Category))
	orm.RegisterDriver("mysql", orm.DR_MySQL)
    orm.RegisterDataBase("default", "mysql", "root:root@/onboarding?charset=utf8")
}



func GetAllCategory() ([]*Category, error) {
	
	o := orm.NewOrm()
	var c []*Category

    name, err := o.QueryTable(Category{}).All(&c)

	if err != orm.ErrNoRows && name > 0 {
		return c, nil
	} else {
		return nil, err
	}
}

func GetById(id int) (*Category, error) {
	o := orm.NewOrm()
	//nid errcnvrt :=strconv.Atoi(id)
	c := &Category{Id: id}
	err := o.Read(c)

	if (err == nil) {
		return c, nil
	} else {
		return nil, errors.New("Data Not Found")
	}
}

func AddCategory(c *Category) (*Category, error) {
	o := orm.NewOrm()
	_, err := o.Insert(c)

	if err != nil {
		log.Print(err)
	}

	return c, err
}

func DeleteCategory(id int) bool {
	o := orm.NewOrm()
	check := o.QueryTable(Category{}).Filter("Id", id).Exist()

	if check {
		num, err := o.Delete(&Category{Id: id})

		if (err == nil && num > 0) {
			return true
		}
	}

	return false
}

func UpdateCategory(c *Category) (int64, bool) {
	o := orm.NewOrm()
	num, err := o.Update(c)
	if  err == nil {
		return num, true
	} else {
		return num, false
	}
}
