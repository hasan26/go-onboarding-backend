package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	Id 			int 
	Title 		string 
	Description string 
	Datetime 	time.Time 
	Archive 	int 
	Priority 	*Priority 
	Status 		*Status 
	Category 	*Category 
	Member 		int
}

type Status struct {

	Id 			int 
	Name 		string 
}

type Priority struct {

	Id 			int 
	Name 		string 

}

func init() {
	
	orm.RegisterModel(new(Todo))
	orm.RegisterDriver("mysql", orm.DR_MySQL)
}

func AddTodo(td *Todo) (*Todo, error) {
	o := orm.NewOrm()
	_, err := o.Insert(td)

	return td, err
}