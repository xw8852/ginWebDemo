package database

import (
	"ginWebDemo/api/util"
	time2 "time"
)

type User struct {
	Id         string
	Phone      string
	Password   string
	UpdateTime time2.Time
}

func UserValidate(phone string) bool {
	d := Default()
	ok := false
	defer d.Db.Close()
	row, err := d.Db.Query("select id from user where phone = ？", phone)
	defer row.Close()
	if !util.Convert(err) {
		var id string
		err := row.Scan(&id)
		if !util.Convert(err) && id != "" {
			ok = true
		}
	}
	return ok
}
func UserLogin(phone string, password string) (User, bool) {
	d := Default()
	var user User
	ok := false
	row, e := d.Db.Query("select * from user where phone = ? and password = ？", phone, password)
	if util.Convert(e) {
		return user, ok
	}
	defer d.Db.Close()
	defer row.Close()
	for row.Next() {
		e := row.Scan(&user.Id, &user.Phone, &user.Password, &user.UpdateTime)
		if util.Convert(e) {
			return user, ok
		}
	}
	ok = user.Id != ""
	return user, ok
}

func RegisterUser(phone string, password string) bool {
	db := Default()
	defer db.Db.Close()
	stmt, err := db.Db.Prepare("INSERT into user (phone,password)VALUES(?,?); ")
	if util.Convert(err) {
		return false
	}
	_, err = stmt.Exec(phone, password)
	if util.Convert(err) {
		return false
	}
	return true
}
