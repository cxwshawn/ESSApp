package main

import (
    "github.com/astaxie/beego/orm"
    _ "ESSApp/models"
    _ "github.com/go-sql-driver/mysql" // import your used driver
)

func init() {
    // // register model
    // orm.RegisterModel(new(User))
    // set default database
    orm.RegisterDataBase("default", "mysql", "root:root@/test?charset=utf8", 30)
}    
func main() {
    //TODO: 命令模式自动建表
    orm.RunCommand()    	   
}