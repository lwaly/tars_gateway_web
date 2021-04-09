package models

import (
	"errors"
	fmt "fmt"
	"strings"

	"github.com/lwaly/tars_gateway_web/server/common"
	"github.com/lwaly/tars_gateway_web/server/models/mymysql"
)

func GatewayHeartInsert(info *GatewayHeart) (code int, err error) {
	conn := mymysql.GetConn()
	if conn == nil {
		err = errors.New("ERR_DATABASE")
		code = common.ERR_DATABASE
		common.Errorf("fail to connect mysql.err=%v", err)
		return
	}

	s := fmt.Sprintf("INSERT INTO %s(`gateway_id`,`status`,`create_time`, `update_time`) values(?, ?, ?, ?, ?, ?, ?)", t_gateway_heart)

	_, err = conn.Exec(s, info.GetGatewayId(), info.GetStatus(), info.GetCreateTime(), info.GetUpdateTime())
	common.Infof("%s %v", s, info)
	if err != nil {
		common.Errorf("fail to exec sql: %s, error: %v", s, err)
		code = common.ERR_DATABASE
		return
	}

	return
}

func GatewayHeartGet(info *GatewayHeart) (code int, err error, res []*GatewayHeart) {
	if mymysql.GetConn() == nil {
		err = errors.New("ERR_DATABASE")
		code = common.ERR_DATABASE
		common.Errorf("fail to connect mysql.err=%v", err)
		return
	}

	s := fmt.Sprintf("SELECT `id`,`gateway_id`,`status`,`create_time`, `update_time` FROM %s", t_gateway_heart)

	values := []interface{}{}
	conditions := []string{}

	if 0 != info.GetStatus() {
		values = append(values, info.GetStatus())
		conditions = append(conditions, "`status`=?")
	}

	if 0 != info.GetGatewayId() {
		values = append(values, info.GetGatewayId())
		conditions = append(conditions, "`gateway_id`=?")
	}

	if 0 < len(conditions) {
		s += " where " + strings.Join(conditions, " AND ")
	}

	rows, err := mymysql.GetConn().Query(s, values...)
	if err != nil {
		common.Errorf("fail to exec sql: %s, error: %v", s, err)
		code = common.ERR_DATABASE
		return
	}
	defer rows.Close()
	for rows.Next() {
		temp := GatewayHeart{}
		if rows.Scan(&temp.Id, &temp.GatewayId, &temp.Status, &temp.CreateTime, &temp.UpdateTime) != nil {
			common.Errorf("Scan数据行失败：err=%v", err)
		}
		res = append(res, &temp)
	}
	return
}
