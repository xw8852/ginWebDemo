package database

import "log"

func InsertSmsCode(phone string, code string) bool {
	db := Default()

	defer db.Db.Close()
	stmt, err := db.Db.Prepare("INSERT into sms_record (phone,msg,status)VALUES(?,?,1); ")
	if (err != nil) {
		log.Fatal(err)
	}
	_, err1 := stmt.Exec(phone, code)
	if (err1 != nil) {
		log.Fatal(err1)
	}
	return true
}
