package controllers

import (
	"go-onboarding-backend/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
	"strconv"
)

// Todo API
type TodoController struct {
	beego.Controller
}

// @Title createdTodo
// @Description create todo
// @Param	body		body 	models.Todo	true		"body for category content"
// @Success 200 {int} models.Todo.Id
// @Failure 403 body is empty
// @router / [post]
func (t *TodoController) Post() {
	var ob models.Todo
	json.Unmarshal(t.Ctx.Input.RequestBody, &ob)
	category, err := models.AddTodo(&ob)

	if err == nil {
		t.Data["json"] = map[string]interface{}{"status": true, "message": "Todo successfully saved.", "result": category}
	} else {
		t.Data["json"] = map[string]interface{}{"status": false, "message": "Todo unsuccessfully saved."}
	}

	t.ServeJson()
}

// @Title getAlltodo
// @Description get all data todo
// @Success 200 {int} models.Todo
// @router / [get]
func (t *TodoController ) GetAll() {
	result, err := models.GetAllTodo()

	if (err == nil) {
		t.Data["json"] = map[string] interface{} {"status": true, "Todo haha": result}
	} else {
		t.Data["json"] = map[string] interface{} {"status": false, "message": "Data Not Found"}
	}

	t.ServeJson()
}

// @Title getTodoByDate
// @Description get all todo by date
// @Param	date		path 	time.Time	true		"The key for staticblock"
// @Success 200 {object} models.Todo
// @Failure 403 :datetime is empty
// @router /:datetime [get]
func (t *TodoController ) Get() {

	tDate := t.Ctx.Input.Param(":id")
	Date, err := time.Parse("2006-01-02 15:04", tDate)

	if (err == nil) {
		result, errs := models.GetByDate(Date)

		if (errs != nil) {
			t.Data["json"] = map[string] interface{} {"status": false,"message": errs.Error()}
		} else {
			t.Data["json"] = result
		}
	} else {
		t.Data["json"] = map[string] interface{} {"status": false,"message": err ,"date":tDate}
	}

	t.ServeJson()
}

// @Title updateTodo
// @Description update the todo
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.Todo
// @Failure 403 :id is not int
// @router /:id [put]
func (t *TodoController ) Put() {

	var ob models.Todo
	tId := t.Ctx.Input.Param(":id")
	Id, err := strconv.Atoi(tId)

	_, err = models.GetTodoById(Id);
	if  err == nil {
		_, success := models.UpdateTodo(&ob)

		if success {
			t.Data["json"] = map[string] interface{} {"status": true,"message": " Todo successfully updated."}
		} else {
			t.Data["json"] = map[string] interface{} {"status": false,"message": "Failed update."}
		}
	} else {
		t.Data["json"] = map[string] interface{} {"status": false,"message": "Todo not found "}
	}

	t.ServeJson()
}


