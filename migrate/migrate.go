package main

import (
	"fmt"
	"goReactLearning/db"
	"goReactLearning/model"
)

//migrationコマンドの実行はプログラム実行開始時に配置したいのでmainパッケージに含める(エントリーポイントに配置)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	//.AutoMigrate()にDBに反映させたいモデル構造を渡す。
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
