package model

import (
	"sync"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

var MySQLPrefix = "MySQL"
var MySQLOnce sync.Once

func MySQLInit() {
	l := g.Log(MySQLPrefix)
	l.SetPrefix(MySQLPrefix)
	//告诉 GoFrame："default-jinli-db(在配置文件配置的) 数据库的所有操作日志，都用 l 这个日志对象来记录"。
	g.DB("default").SetLogger(l) //在配置文件配置
}

func GetDB(name ...string) (db gdb.DB) {
	MySQLOnce.Do(MySQLInit)
	return g.DB(name...)
}
