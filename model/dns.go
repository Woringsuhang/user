package model

import (
	"fmt"
)

func Dns(username, password, host string, port int, library string) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, library)
	fmt.Println(dsn)
	return dsn
}
