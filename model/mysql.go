package model

import (
	"fmt"
	"github.com/Woringsuhang/user/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() error {
	var err error
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.ConfigAll.Mysql.Username, global.ConfigAll.Mysql.Password, global.ConfigAll.Mysql.Host,
		global.ConfigAll.Mysql.Port, global.ConfigAll.Mysql.Library)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func Tx(txFu func(tx *gorm.DB) error) {
	var err error

	tx := DB.Begin()
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

}
