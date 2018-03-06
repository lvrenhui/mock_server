package controllers

import (
	"encoding/json"
	m "mock_server/models"

	"github.com/astaxie/beego"
)

//MockController handle mock request
type MockController struct {
	beego.Controller
}

func (c *MockController) Get() {
	id, _ := c.GetInt("id")
	var data m.MockData
	m.GetOne(id, &data)
	c.Data["json"] = &data
	c.ServeJSON()
}

//List ...
func (c *MockController) List() {
	// ids := []int{1, 2, 3}
	var (
		total    = 0
		size     = 10
		dataList []m.MockData
	)
	page, _ := c.GetInt("p", 1)
	m.GetList(page, size, &total, &dataList)
	// if err1 != nil || err2 != nil {
	// 	beego.Error(err.Error())
	// 	c.Rsp(false, err.Error())
	// 	return
	// }
	c.Data["Page"] = m.PageUtil(total, page, size, dataList)
	// c.ServeJSON()
	c.TplName = "list.html"
}

// AddPage ...
func (c *MockController) AddPage() {
	c.TplName = "add.html"
}

//Add mock data
func (c *MockController) Add() {
	data := m.MockData{}
	// beego.Debug(c.Input())
	if err := c.ParseForm(&data); err != nil {
		//handle error
		c.Rsp(false, err.Error())
		return
	}
	_, err := m.Add(&data)
	if err == nil {
		c.Rsp(true, "ok")
		// fmt.Println(users)
	} else {
		c.Rsp(false, err.Error())
	}
}

//Del ...
func (c *MockController) Del() {
	// beego.Debug(data)
	id, _ := c.GetInt("id")
	_, err := m.Delete(id)
	if err == nil {
		c.Rsp(true, "ok")
		// fmt.Println(users)
	} else {
		c.Rsp(false, err.Error())
	}
}

//UpdatePage ..
func (c *MockController) UpdatePage() {
	id, _ := c.GetInt("id")
	c.Data["id"] = id
	var model m.MockData
	m.GetOne(id, &model)
	c.Data["model"] = model
	c.TplName = "add.html"
}

//Update ...
func (c *MockController) Update() {
	// beego.Debug(data)
	id, _ := c.GetInt("id")
	data := m.MockData{}
	// beego.Debug(c.Input())
	if err := c.ParseForm(&data); err != nil {
		//handle error
		c.Rsp(false, err.Error())
		return
	}
	_, err := m.Update(id, data)
	if err == nil {
		c.Rsp(true, "ok")
		// fmt.Println(users)
	} else {
		c.Rsp(false, err.Error())
	}
}

//UpdateCheck ...
func (c *MockController) UpdateCheck() {
	// beego.Debug(data)
	id, _ := c.GetInt("id")
	checked, _ := c.GetInt("check")
	_, err := m.UpdateCheck(id, checked)
	if err == nil {
		c.Rsp(true, "ok")
		// fmt.Println(users)
	} else {
		c.Rsp(false, err.Error())
	}
}

//Rsp ,common rsp
func (c *MockController) Rsp(status bool, str string) {
	c.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	c.ServeJSON()
}

//Mock ...
func (c *MockController) Mock() {
	uri := c.Ctx.Request.RequestURI
	data, err := m.GetMock(uri)
	if err == nil {
		var obj map[string]interface{}
		if err := json.Unmarshal([]byte(data), &obj); err == nil {
			c.Data["json"] = &obj
			c.ServeJSON()
		} else {
			c.Rsp(false, "json format error!")
		}

	} else {
		c.Rsp(false, "get from db error")
	}

}
