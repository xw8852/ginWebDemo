package database

import (
	"ginWebDemo/api/util"
	"fmt"
)

type User struct {
	Id         string
	Phone      string
	Password   string
	UpdateTime string
}

func UserCount() int {
	count := 0
	d := Default()
	defer d.Db.Close()
	row, err := d.Db.Query("select count(id) from user")
	fmt.Println(err)
	defer row.Close()
	if !util.Convert(err) {
		if row.Next() {
			err := row.Scan(&count)
			fmt.Println(err)
		}
	}
	return count
}
func UserPhoneValidate(phone string) bool {
	d := Default()
	defer d.Db.Close()
	ok := false
	row, err := d.Db.Query("select id from user where phone = ?", phone)
	defer row.Close()
	if !util.Convert(err) {
		var id string
		if row.Next() {
			err := row.Scan(&id)
			if !util.Convert(err) && id != "" {
				ok = true
			}
		}
	}
	return ok
}
func UserLoginNameValidate(logonName string) bool {
	d := Default()
	defer d.Db.Close()
	ok := false
	row, err := d.Db.Query("select id from user where logonName = ?", logonName)
	defer row.Close()
	if !util.Convert(err) {
		var id string
		if row.Next() {
			err := row.Scan(&id)
			if !util.Convert(err) && id != "" {
				ok = true
			}
		}
	}
	return ok
}
func UserGet(id string) User {
	d := Default()
	var user User
	row, e := d.Db.Query("select id,phone,updateTime from user where id = ?  ", id)
	if util.Convert(e) {
		return user
	}
	defer d.Db.Close()
	defer row.Close()
	for row.Next() {
		e := row.Scan(&user.Id, &user.Phone, &user.UpdateTime)
		if util.Convert(e) {
			return user
		}
	}

	return user
}
func UserLoginByName(logonName string, password string) (User, bool) {
	d := Default()
	var user User
	ok := false
	row, e := d.Db.Query("select id,phone,updateTime from user where logonName = ? and password = ? ", logonName, password)
	if util.Convert(e) {
		return user, ok
	}
	defer d.Db.Close()
	defer row.Close()
	for row.Next() {
		e := row.Scan(&user.Id, &user.Phone, &user.UpdateTime)
		if util.Convert(e) {
			return user, ok
		}
	}
	ok = user.Id != ""
	return user, ok
}
func UserLogin(phone string, password string) (User, bool) {
	d := Default()
	var user User
	ok := false
	row, e := d.Db.Query("select id,phone,updateTime from user where phone = ? and password = ? ", phone, password)
	if util.Convert(e) {
		return user, ok
	}
	defer d.Db.Close()
	defer row.Close()
	for row.Next() {
		e := row.Scan(&user.Id, &user.Phone, &user.UpdateTime)
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
func RegisterUserByName(name string, password string) bool {
	db := Default()
	defer db.Db.Close()
	stmt, err := db.Db.Prepare("INSERT into user (logonName,password)VALUES(?,?); ")
	if util.Convert(err) {
		return false
	}
	_, err = stmt.Exec(name, password)
	if util.Convert(err) {
		return false
	}
	return true
}
