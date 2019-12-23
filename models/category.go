package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
)

func FindCategoryByName(name string) (error, *Category) {
	o := orm.NewOrm()

	c := Category{Title: name}
	err := o.Read(&c, "title") // 找某条记录，第二个参数是根据哪个索引查找
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}
		return err, nil
	}
	return nil, &c
}

func FindCategoryByNameVersion2(name string) (error, *Category) {
	o := orm.NewOrm()

	c := Category{Title: name}
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(&c) // 注意参数是：指针
	if err != nil {
		return err, nil
	}
	return nil, &c
}

func AddCategory(name string) error {
	o := orm.NewOrm()

	c := Category{Title: name}
	//尝试从数据库读取，不存在的话就创建一个
	//默认必须传入一个参数作为条件字段，同时也支持多个参数多个条件字段
	var err error
	var created bool
	var id int64
	if created, id, err = o.ReadOrCreate(&c, "title"); err == nil {
		fmt.Println("internal", err)
		if created {
			fmt.Println("New Insert an object. Id:", id)
			return nil
		} else {
			fmt.Println("Get an object. Id:", id)
			return errors.New("has exist")
		}
	}
	fmt.Println("AddCategory err,", err)
	return err
}

func GetAllCategory() (error, []*Category) {
	engine := orm.NewOrm()

	c := make([]*Category, 0)
	qs := engine.QueryTable("category")
	_, err := qs.All(&c)
	if err != nil {
		return err, nil
	}
	return nil, c
}

func DelCategoryById(id int) error {
	engine := orm.NewOrm()

	c := Category{Id: int64(id)}
	if _, err := engine.Delete(&c); err != nil {
		return err
	}
	return nil
}
