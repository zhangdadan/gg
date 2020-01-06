package main

import (
	"dakesolo.com/gg/share"
	"log"
	"net/http"
)

func main() {

	//初始化路径
	share.InitPath()

	//初始化日志
	share.InitPath()

	//初始化数据库
	share.InitDB()

	//初始化OTS
	share.InitOTS()

	//初始化OTS
	share.InitLog()

	//配置路由
	mux := http.NewServeMux()
	route(mux)

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}