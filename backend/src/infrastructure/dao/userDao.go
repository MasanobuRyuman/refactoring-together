package dao

import (
	"everybodys_refactoring/src/domain/entity"
	"github.com/guregu/dynamo"
	"time"
)

type userTable struct {
	table dynamo.Table
}

func NewUserTable(table dynamo.Table) *userTable {
	return &userTable{
		table: table,
	}
}

func (db userTable) Add(id string, name string) (err error) {
	err = db.table.Put(&entity.User{Id: id, Name: name, CreateTime: time.Now(), UpdateTime: time.Now()}).Run()
	return
}

func (db userTable) FindAll() (results []entity.User, err error) {
	err = db.table.Scan().All(&results)
	return
}

func (db userTable) FindById(id string) (result entity.User, err error) {
	err = db.table.Get("Id", id).One(&result)
	return
}

func (db userTable) Update(id string,name string)(err error){
  err = db.table.Update("Id", id).Set("Name",name).Run()
	return
}
