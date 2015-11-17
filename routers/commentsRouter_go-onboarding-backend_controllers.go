package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["go-onboarding-backend/controllers:CategoryController"] = append(beego.GlobalControllerRouter["go-onboarding-backend/controllers:CategoryController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-onboarding-backend/controllers:CategoryController"] = append(beego.GlobalControllerRouter["go-onboarding-backend/controllers:CategoryController"],
		beego.ControllerComments{
			"Get",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-onboarding-backend/controllers:CategoryController"] = append(beego.GlobalControllerRouter["go-onboarding-backend/controllers:CategoryController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-onboarding-backend/controllers:CategoryController"] = append(beego.GlobalControllerRouter["go-onboarding-backend/controllers:CategoryController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["go-onboarding-backend/controllers:CategoryController"] = append(beego.GlobalControllerRouter["go-onboarding-backend/controllers:CategoryController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["go-onboarding-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["go-onboarding-backend/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-onboarding-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["go-onboarding-backend/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-onboarding-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["go-onboarding-backend/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-onboarding-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["go-onboarding-backend/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["go-onboarding-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["go-onboarding-backend/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["go-onboarding-backend/controllers:TodoController"] = append(beego.GlobalControllerRouter["go-onboarding-backend/controllers:TodoController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-onboarding-backend/controllers:TodoController"] = append(beego.GlobalControllerRouter["go-onboarding-backend/controllers:TodoController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-onboarding-backend/controllers:TodoController"] = append(beego.GlobalControllerRouter["go-onboarding-backend/controllers:TodoController"],
		beego.ControllerComments{
			"Get",
			`/:datetime`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-onboarding-backend/controllers:TodoController"] = append(beego.GlobalControllerRouter["go-onboarding-backend/controllers:TodoController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["go-onboarding-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["go-onboarding-backend/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["go-onboarding-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["go-onboarding-backend/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-onboarding-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["go-onboarding-backend/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-onboarding-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["go-onboarding-backend/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["go-onboarding-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["go-onboarding-backend/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["go-onboarding-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["go-onboarding-backend/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["go-onboarding-backend/controllers:UserController"] = append(beego.GlobalControllerRouter["go-onboarding-backend/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

}
