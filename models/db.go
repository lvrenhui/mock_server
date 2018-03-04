package models

import (
	"database/sql"
	"time"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	"github.com/beego/admin/src/models"
)

func GetList(page int, size int, total *int, dataList *[]MockData) {
	models.Connect()
	o := orm.NewOrm()
	o.Raw("select count(*) from mockdata").QueryRow(total)
	o.Raw("SELECT * FROM mockdata order by id desc limit ?,?", (page-1)*size, size).QueryRows(dataList)

}

func GetOne(id int, data *MockData) {
	models.Connect()
	o := orm.NewOrm()
	o.Raw("SELECT * FROM mockdata where id=?", id).QueryRow(data)

}

func Add(m *MockData) (sql.Result, error) {
	models.Connect()
	// beego.Debug(data)
	timeNow := time.Now()
	o := orm.NewOrm()
	return o.Raw("insert into mockdata values(null,?,?,?,?,?,?)", m.Path, m.Data, m.Tip, 1, timeNow, timeNow).Exec()
}

func Update(id int, data MockData) (sql.Result, error) {
	// beego.Debug(data)
	models.Connect()
	o := orm.NewOrm()
	timeNow := time.Now()
	return o.Raw("update mockdata set path=?,data=?,tip=?,update_time=? where id=?", data.Path, data.Data, data.Tip, timeNow, id).Exec()

}

func UpdateCheck(id int, checked int) (sql.Result, error) {
	// beego.Debug(data)
	models.Connect()
	o := orm.NewOrm()
	beego.Debug(checked)
	beego.Debug(id)
	return o.Raw("update mockdata set `check`=? where id=?", checked, id).Exec()

}

func Delete(id int) (sql.Result, error) {
	models.Connect()
	// beego.Debug(data)
	o := orm.NewOrm()
	return o.Raw("delete from mockdata where id=?", id).Exec()
}

// GetMock ...
func GetMock(path string) (string, error) {
	models.Connect()
	// beego.Debug(data)
	o := orm.NewOrm()
	var data string
	sql := "SELECT data FROM mockdata where `check`=1 and path='" + path + "' order by update_time desc limit 1"
	err := o.Raw(sql).QueryRow(&data)

	beego.Debug(sql)
	beego.Debug(path)
	beego.Debug(data)

	return data, err

}
