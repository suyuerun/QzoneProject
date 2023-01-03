package database

import (
	"fmt"
	"gin/config"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type SqlOpera struct {
}

var Database *sqlx.DB

func init() {
	dataSourceName := config.GlobalConfig.MysqlUser +
		":" + config.GlobalConfig.MysqlPwd +
		"@tcp(" + config.GlobalConfig.MysqlHost +
		":" + strconv.Itoa(config.GlobalConfig.MysqlPort) +
		")/" + config.GlobalConfig.MysqlDBName
	var err error
	Database, err = sqlx.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
	}
	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")

	//defer Database.Close()
}
