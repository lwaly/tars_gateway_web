package controllers

import (
	"time"

	"github.com/kataras/iris"
	"github.com/lwaly/tars_gateway_web/server/common"
	"github.com/lwaly/tars_gateway_web/server/models"
)

func ServerListGet(ctx iris.Context) {
	info := models.GatewayInfo{Ip: []byte("test"), Name: []byte("test"), Area: []byte("test"),
		Status: models.VALID, CreateTime: time.Now().Unix(), UpdateTime: time.Now().Unix()}

	if code, err := models.GatewayInfoInsert(&info); common.OK != code {
		common.Errorf("fail to InsertMaterial%v", err)
	}

	return
}
