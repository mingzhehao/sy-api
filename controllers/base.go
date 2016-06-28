package controllers

import (
        "github.com/astaxie/beego"
)

type BaseController struct {
        beego.Controller
}

func  JsonFormat(retcode int,retmsg string,retdata interface{}) (json map[string]interface{} ){
	if(retcode == 0) {
	 	json = map[string]interface{}{
			"retcode": retcode,
			"retdata": retdata,
			"retmsg": retmsg,
		} 
	} else {
	 	json = map[string]interface{}{
			"retcode": retcode,
			"retmsg": retmsg,
		} 
	}

	return json
}
