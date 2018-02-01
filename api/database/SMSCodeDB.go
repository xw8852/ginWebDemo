package database

import (
	"ginWebDemo/api/util"
)

func InsertSmsCode(phone string, code string) bool {
	db := Default()
	defer db.Db.Close()
	stmt, err := db.Db.Prepare("INSERT into sms_record (phone,msg,status)VALUES(?,?,1); ")
	if util.Convert(err) {
		return false
	}
	_, err1 := stmt.Exec(phone, code)
	if util.Convert(err1) {
		return false
	}
	return true
}
