package models

import (
	"database/sql"
	"github.com/astaxie/beedb"
	_ "github.com/ziutek/mymysql/godrv"
	"time"
)

type Blog struct {
	Id int `PK`
	Title string
	Content string
	Created time.Time
}

func GetLink() beedb.Model {
	beedb.OnDebug = true
	db, err := sql.Open("mymysql", "test/root/123456")
	if err != nil {
		panic(err)
	}

	orm := beedb.New(db)

	return orm
}

func GetAll() (blogs []Blog) {
	db := GetLink()
	db.FindAll(&blogs)
	
	return
}

func GetBlog(id int) (blog Blog) {
	db := GetLink()
	db.Where("id=?", id).Find(&blog)
	return
}

func SaveBlog(blog Blog) {
	db := GetLink()
	db.Save(&blog)
	return
}

func DelBlog(blog Blog) {
	db := GetLink()
	db.Delete(&blog)
	return
}