package controllers

import (
	"go-onboarding-backend/models"
	"encoding/json"
	"strconv"
	"github.com/astaxie/beego"
)

// Category API
type CategoryController struct {
	beego.Controller
}

// @Title getAllCategory
// @Description get all data category
// @Success 200 {int} models.Category
// @router / [get]
func (c *CategoryController ) GetAll() {
	result, err := models.GetAllCategory()

	if (err == nil) {
		c.Data["json"] = map[string] interface{} {"status": true, "Category": result}
	} else {
		c.Data["json"] = map[string] interface{} {"status": false, "message": "No Data Found"}
	}

	c.ServeJson()
}

// @Title Get
// @Description get category by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Category
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CategoryController ) Get() {

	cId := c.Ctx.Input.Param(":id")
	Id, err := strconv.Atoi(cId)

	if (err == nil) {
		result, errs := models.GetById(Id)

		if (errs != nil) {
			c.Data["json"] = map[string] interface{} {"status": false,"message": errs.Error()}
		} else {
			c.Data["json"] = result
		}
	} else {
		c.Data["json"] = map[string] interface{} {"status": false,"message": err ,"id":cId}
	}

	c.ServeJson()
}

// @Title createCategory
// @Description create category
// @Param	body		body 	models.Categori	true		"body for category content"
// @Success 200 {int} models.Category.Id
// @Failure 403 body is empty
// @router / [post]
func (c *CategoryController ) Post() {

	var catbody models.Category
	json.Unmarshal(c.Ctx.Input.RequestBody, &catbody)
	cat, err := models.AddCategory(&catbody)

	if err == nil {
		c.Data["json"] = map[string] interface{} {"status": true,"message": "Category successfully saved.","result": cat}
	} else {
		c.Data["json"] = map[string] interface{} {"status": false,"message": "Category unsuccessfully saved."}
	}

	c.ServeJson()
}


// @Title deleteCategory
// @Description delete the category
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:id [delete]
func (c *CategoryController )  Delete() {
	cId := c.Ctx.Input.Param(":id")
	Id, err := strconv.Atoi(cId)

	if (err == nil) {
		result := models.DeleteCategory(Id)

		if (result) {
			c.Data["json"] = map[string] interface{} {"status": true, "message": "Delete Successfully"}
		} else {
			c.Data["json"] = map[string] interface{} {"status": false, "message": "Delete Failed"}
		}
	} else {
		c.Data["json"] = map[string] interface{} {"status": false,"message": err}
	}

	c.ServeJson()
}

// @Title updateCategory
// @Description update the category
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.Category
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CategoryController ) Put() {

	var ob models.Category
	cId := c.Ctx.Input.Param(":id")
	Id, err := strconv.Atoi(cId)

	_, err = models.GetById(Id);
	if  err == nil {
		_, success := models.UpdateCategory(&ob)

		if success {
			c.Data["json"] = map[string] interface{} {"status": true,"message": " Category successfully updated."}
		} else {
			c.Data["json"] = map[string] interface{} {"status": false,"message": "Failed update."}
		}
	} else {
		c.Data["json"] = map[string] interface{} {"status": false,"message": "Category not found "}
	}

	c.ServeJson()
}





