package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"myBeego/models"
	"strconv"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {

	op := this.Input().Get("op")
	switch op {

	case "add":
		categoryName := this.Input().Get("categoryName")
		fmt.Println("main CategoryController get,categoryName", categoryName)
		if len(categoryName) == 0 {
			break
		}
		err := models.AddCategory(categoryName)
		if err != nil {
			logs.Info("models.AddCategory error:", err)
		}
		//this.Redirect("/category", 301)  // 可以 不跳转，因为下面已经有获取所有数据，然后展示了
		//return
	case "del":
		idStr := this.Input().Get("id")
		if len(idStr) == 0 {
			break
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			break
		}

		err = models.DelCategoryById(id)
		if err != nil {
			logs.Info("models.AddCategory error:", err)
		}
	}

	this.Data["isCategory"] = true
	var err error
	err, this.Data["CategoryList"] = models.GetAllCategory()
	if err != nil {
		logs.Error("GetAllCategory err:", err)
	}
	this.TplName = "category.html"

}
