package controllers

import (
	"go-onboarding-backend/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

// Categori API
type TodoController struct {
	beego.Controller
}

// @Title created new todo
// @Description create todo
// @Param	body		body 	models.Todo	true		"body for todo content"
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





