package main

import (
	"fmt"
	"wallet_real/config"
	"wallet_real/pkg/pg"
)

func main() {

	config.SetEnv()

	db, err := pg.ConnectDB(config.AppConfig{
		Host:     config.LoadEnv().Host,
		Port:     config.LoadEnv().Port,
		Username: config.LoadEnv().Username,
		Password: config.LoadEnv().Password,
		DBname:   config.LoadEnv().DBname,
	})
	// error handling
	if err != nil {
		fmt.Println("Đã có lỗi xảy ra: ", err)
		return
	}
	fmt.Println(db)
}
