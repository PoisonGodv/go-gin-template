package main

import (
	"test_wxlogin/router"
	"test_wxlogin/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	utils.InitRedis()
	r := router.Router()
	r.Run(":9000")
}
