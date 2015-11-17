package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	Id         		int    `orm:"column(id);auto"`
	Title       	string `orm:"column(title);size(50)"`
	Description		string `orm:"column(description);size(300)"`
	Datetime		time.Time `orm:"column(datetime);auto_now_add;type(datetime)"`
	Archive			int    
	Priority    	*Priority `orm:"rel(fk);column(id_priority)"`
	Status			*Status `orm:"rel(fk);column(id_status)"`
	Category		*Category `orm:"rel(fk);column(id_category)"`
	User			int    
}

type Status struct {
	Id 			int    `orm:"column(id);auto"` 
	Name 		string `orm:"column(title);size(50)"` 
}

type Priority struct {
	Id 			int    `orm:"column(id);auto"`  
	Name 		string `orm:"column(title);size(50)"` 

}

func init() {
	
	orm.RegisterModel(new(Todo))
	orm.RegisterModel(new(Status))
	orm.RegisterModel(new(Priority))
	orm.RegisterDriver("mysql", orm.DR_MySQL)
}

func AddTodo(td *Todo) (*Todo, error) {
	o := orm.NewOrm()
	_, err := o.Insert(td)

	return td, err
}

func GetAllTodo() ([]*Todo, error) {
	
	o := orm.NewOrm()
	var t []*Todo

    name, err := o.QueryTable(Todo{}).RelatedSel().All(&t)

	if err != orm.ErrNoRows && name > 0 {
		return t, nil
	} else {
		return t, nil
	}
}

func GetByDate(datetime time.Time) (*Todo, error) {
	
	o := orm.NewOrm()
	t := &Todo{Datetime: datetime}

	err := o.QueryTable("todo").Filter("Datetime", datetime).RelatedSel().One(t);

	if  err == nil {
		return t, nil
	} else {
		return nil,err
	}
}

func Archrive(t *Todo) bool {
	o := orm.NewOrm()
	num, err := o.Update(t, "archive")
	if (err == nil && num > 0) {
		return true
	}

	return false
}

func UpdateTodo(t *Todo) (int64, bool) {
	o := orm.NewOrm()
	num, err := o.Update(t)

	if  err == nil {
		return num, true
	} else {
		return num, false
	}
}

func GetTodoById(id int) (*Todo, error) {
	o := orm.NewOrm()
	//nid errcnvrt :=strconv.Atoi(id)
	t := &Todo{Id: id}
	err := o.Read(t)

	if (err == nil) {
		return t, nil
	} else {
		return nil, err
	}
}
