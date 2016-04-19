package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["sy-project/sy-api/controllers:AppController"] = append(beego.GlobalControllerRouter["sy-project/sy-api/controllers:AppController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["sy-project/sy-api/controllers:AppController"] = append(beego.GlobalControllerRouter["sy-project/sy-api/controllers:AppController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["sy-project/sy-api/controllers:AppController"] = append(beego.GlobalControllerRouter["sy-project/sy-api/controllers:AppController"],
		beego.ControllerComments{
			"GetJson",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["sy-project/sy-api/controllers:AppController"] = append(beego.GlobalControllerRouter["sy-project/sy-api/controllers:AppController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["sy-project/sy-api/controllers:AppController"] = append(beego.GlobalControllerRouter["sy-project/sy-api/controllers:AppController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["sy-project/sy-api/controllers:AppController"] = append(beego.GlobalControllerRouter["sy-project/sy-api/controllers:AppController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["sy-project/sy-api/controllers:UserController"] = append(beego.GlobalControllerRouter["sy-project/sy-api/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["sy-project/sy-api/controllers:UserController"] = append(beego.GlobalControllerRouter["sy-project/sy-api/controllers:UserController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["sy-project/sy-api/controllers:UserController"] = append(beego.GlobalControllerRouter["sy-project/sy-api/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["sy-project/sy-api/controllers:UserController"] = append(beego.GlobalControllerRouter["sy-project/sy-api/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["sy-project/sy-api/controllers:UserController"] = append(beego.GlobalControllerRouter["sy-project/sy-api/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
