package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"sy-project/sy-api/models"

	"github.com/astaxie/beego"
)

// oprations for App
type AppController struct {
	beego.Controller
}

func (c *AppController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetJson", c.GetJson)
}

// @Title Post
// @Description create App
// @Param	body		body 	models.App	true		"body for App content"
// @Success 201 {int} models.App
// @Failure 403 body is empty
// @router / [post]
func (c *AppController) Post() {
	var v models.App
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddApp(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Get
// @Description get App by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.App
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AppController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetAppById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title GetJson
// @Description get App by id (2维json)
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.App
// @Failure 403 :id is empty
// @router /getjson/:id [get]
func (c *AppController) GetJson() {
	/* :id 与路由定义的:id 要统一名称 */
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetAppById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		/** 自定义格式输出 多维json -- start **/
		c.Data["json"] = map[string]interface{}{
			"retcode": 0,
			"retdata": map[string]interface{}{
				"AppCode":    v.AppCode,
				"AppName":    v.AppName,
				"CreateDate": v.CreateDate},
			"retmsg": "success"}
		/** 自定义格式输出 多维json -- end **/
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get App
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.App
// @Failure 403
// @router / [get]
func (c *AppController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllApp(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the App
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.App	true		"body for App content"
// @Success 200 {object} models.App
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AppController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.App{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateAppById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the App
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AppController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteApp(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
