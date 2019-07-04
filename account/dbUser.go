package account

import (
	"server/data"

	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2/bson"
)

// 账户信息
type DBAcc struct {
	ID     int64  `bson:"uid"`
	ACNAME string `bson:"name"`
	PWD    string `bson:"password"`
}

type DBUser struct {
	ID     int64  `bson:"uid"`
	EXP    int64  `bson:"exp"`
	NAME   string `bson:"name"`
	IC     int    `bson:"icon"`
	TOTAL  int64  `bson:"playcount"`
	WIN    int64  `bson:"playwin"`
	OUT    int64  `bson:"playout"`
	CREATE int64  `bson:"create"`
	HONOR  int64  `bson:"honor"`
	MONEY  int64  `bson:"money"`
	GLOD   int64  `bson:"gold"`
	TITLE  int    `bson:"title"`
	STATYS int    `bson:"status"`
}

var (
	dbName  = "runmouse"
	tabAcc  = "admin"
	tabUser = "user"
)

// 从数据库中查找账号
func FindDBAcc(name string, id int64) *DBAcc {
	acc := new(DBAcc)

	var err error
	if id <= 0 {
		err = data.Find(dbName, tabAcc, bson.M{"name": name}, &acc)
	} else {
		err = data.Find(dbName, tabAcc, bson.M{"uid": id}, &acc)
	}

	if err != nil {
		log.Error("user.go - findDBAcc error : %v", err)
		return nil
	}

	return acc
}

// 从数据库中查找用户
func FindDBUser(id int64) *DBUser {
	u := new(DBUser)

	err := data.Find(dbName, tabUser, bson.M{"uid": id}, &u)

	if err != nil {
		log.Error("user.go - findDBUser error : %v", err)

		return nil
	}

	return u
}

// 更新账户表
func UpdateAccDB(acc *DBAcc) {
	if acc == nil {
		return
	}

	data.Update(dbName, tabAcc, bson.M{"uid": acc.ID}, acc)

}

// 更新用户数据库表
func UpdateUserDB(user *DBUser) {
	if user == nil {
		return
	}

	data.Update(dbName, tabUser, bson.M{"uid": user.ID}, user)
}

// 新增账户表数据
func InsertAccDB(acc *DBAcc) {
	if acc == nil {
		return
	}

	data.Insert(dbName, tabAcc, acc)
}

// 新增用户表数据
func InsertUserDB(user *DBUser) {
	if user == nil {
		return
	}

	data.Insert(dbName, tabUser, user)
}
