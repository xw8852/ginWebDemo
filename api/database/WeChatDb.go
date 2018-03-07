package database

import "ginWebDemo/api/util"

func InsertWeChatRelation(userId string,openid string) bool{
	db := Default()
	defer db.Db.Close()
	stmt, err := db.Db.Prepare("INSERT into user_wechat (userid,wechatId)VALUES(?,?); ")
	if util.Convert(err) {
		return false
	}
	_, err = stmt.Exec(userId, openid)
	if util.Convert(err) {
		return false
	}
	return true
}
func UserGetByWeChat(openid string) (User, bool) {
	d := Default()
	var user User
	ok := false
	row, e := d.Db.Query("select u.id,u.phone,u.updateTime from user u,user_wechat w where u.id = w.userid and w.wechatId = ？", openid)
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
