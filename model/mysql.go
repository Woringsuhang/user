package model

import (
	"github.com/Woringsuhang/user/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() {
	var err error
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := Dns(global.ConfigAll.Mysql.Username, global.ConfigAll.Mysql.Password, global.ConfigAll.Mysql.Host,
		global.ConfigAll.Mysql.Port, global.ConfigAll.Mysql.Library)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Dbs, _ := DB.DB()
	defer Dbs.Close()
}

func Tx(txFu func(tx *gorm.DB) error) {
	var err error

	tx := DB.Begin()
	err = txFu(tx)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

}
