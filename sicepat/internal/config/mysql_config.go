package config

import "fmt"

type MysqlConf struct {
	Host         string `json:"host"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseName string `json:"databaseName"`
}

func (a *AppConfig) DatabaseServerName() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True",
		a.Mysql.Username, a.Mysql.Password, a.Mysql.Host, a.Mysql.DatabaseName)
}
