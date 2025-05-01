package main

import (
	"ginblog/model"
	"ginblog/routes"
)

func main() {
	model.InitDb()      // 初始化数据库
	routes.InitRouter() // 初始化路由
}
