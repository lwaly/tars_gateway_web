package controllers

import "github.com/lwaly/tars_gateway_web/server/models/mymysql"

func InitLogic() {
	mymysql.InitMysql()
}
