package controllers

import (
	"bupt_tour/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

//  PlaceController operations for Place
type PlaceController struct {
	beego.Controller
}

// URLMapping ...
func (c *PlaceController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Place
// @Param	body		body	models.Place	true		"body for Place content"
// @Success 201 {int} models.Place
// @Failure 403 body is empty
// @router / [post]
func (c *PlaceController) Post() {
	var v models.Place
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		c.Data["json"] = err.Error()
		c.Abort("400")
	}

	if _, err := models.AddPlace(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
		c.Abort("500")
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Place by id
// @Param	id		path	string	true		"The key for staticblock"
// @Success 200 {object} models.Place
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PlaceController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		c.Data["json"] = err.Error()
		c.Abort("400")
	}
	v, err := models.GetPlaceById(id)
	if err != nil {
		c.Data["json"] = err.Error()
		c.Abort("500")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Place
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Place
// @Failure 403
// @router / [get]
func (c *PlaceController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

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
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllPlace(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
		c.Abort("500")
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Place
// @Param	id		path	string	true		"The id you want to update"
// @Param	body		body	models.Place	true		"body for Place content"
// @Success 200 {object} models.Place
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PlaceController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Place{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdatePlaceById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
		c.Abort("500")
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Place
// @Param	id		path	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PlaceController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeletePlace(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
		c.Abort("500")
	}
	c.ServeJSON()
}

//Match
// @Title Match
// @Description 匹配最近的地址信息
// @Param longitude formData number true "经度"
// @Param latitude formData number true "纬度"
// @Success 200 {object} models.Place
// @Failure 400 {string} 输入错误！
// @router /match [post]
func (c *PlaceController) Match() {
	longitude, err := c.GetFloat("longitude")
	if err != nil {
		c.Data["json"] = err.Error()
		c.Abort("400")
		return
	}
	latitude, err := c.GetFloat("latitude")
	if err != nil {
		c.Data["json"] = err.Error()
		c.Abort("400")
		return
	}
	if v, err := models.MatchPlace(longitude, latitude); err == nil {
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
		c.Abort("500")
	}
	c.ServeJSON()
}
