package models

import (
	"errors"
	fmt "fmt"
	"strings"

	"github.com/lwaly/tars_gateway_web/server/models/mymysql"

	"github.com/lwaly/tars_gateway_web/server/common"
)

func GatewayInfoInsert(info *GatewayInfo) (code int, err error) {
	conn := mymysql.GetConn()
	if conn == nil {
		err = errors.New("ERR_DATABASE")
		code = common.ERR_DATABASE
		common.Errorf("fail to connect mysql.err=%v", err)
		return
	}

	s := fmt.Sprintf("INSERT INTO %s(`ip`,`name`,`area`,`status`,`create_time`, `update_time`) values(?, ?, ?, ?, ?, ?, ?)", t_gateway_info)

	_, err = conn.Exec(s, info.GetIp(), info.GetName(), info.GetArea(), info.GetStatus(), info.GetCreateTime(), info.GetUpdateTime())
	common.Infof("%s %v", s, info)
	if err != nil {
		common.Errorf("fail to exec sql: %s, error: %v", s, err)
		code = common.ERR_DATABASE
		return
	}

	return
}

func GatewayInfoGet(info *GatewayInfo) (code int, err error, res []*GatewayInfo) {
	if mymysql.GetConn() == nil {
		err = errors.New("ERR_DATABASE")
		code = common.ERR_DATABASE
		common.Errorf("fail to connect mysql.err=%v", err)
		return
	}

	s := fmt.Sprintf("SELECT `id`,`ip`,`name`,`area`,`status`,`create_time`, `update_time` FROM %s", t_gateway_info)

	values := []interface{}{}
	conditions := []string{}

	if 0 != info.GetStatus() {
		values = append(values, info.GetStatus())
		conditions = append(conditions, "`status`=?")
	}

	if 0 != len(info.GetIp()) {
		values = append(values, info.GetIp())
		conditions = append(conditions, "`ip`=?")
	}

	if 0 != len(info.GetName()) {
		values = append(values, info.GetName())
		conditions = append(conditions, "`name`=?")
	}

	if 0 != len(info.GetArea()) {
		values = append(values, info.GetArea())
		conditions = append(conditions, "`area`=?")
	}

	if 0 < len(conditions) {
		s += " where " + strings.Join(conditions, " AND ")
	}
	common.Infof("%s %v", s, values)
	rows, err := mymysql.GetConn().Query(s, values...)
	if err != nil {
		common.Errorf("fail to exec sql: %s, error: %v", s, err)
		code = common.ERR_DATABASE
		return
	}
	defer rows.Close()
	for rows.Next() {
		temp := GatewayInfo{}
		if rows.Scan(&temp.Id, &temp.Ip, &temp.Name, &temp.Area, &temp.Status, &temp.CreateTime, &temp.UpdateTime) != nil {
			common.Errorf("Scan数据行失败：err=%v", err)
		}
		res = append(res, &temp)
	}
	return
}

func GatewayInfoUpdate(info *GatewayInfo) (code int, err error) {
	code = common.OK

	if nil == mymysql.GetConn() {
		err = errors.New("ERR_DATABASE")
		code = common.ERR_DATABASE
		common.Errorf("fail to get mysql connect, error: %v", err)
		return
	}

	sql := fmt.Sprintf("update %s set `status`=?, `update_time`=? where `id`=?", t_gateway_info)

	_, err = mymysql.GetConn().Exec(sql, info.GetStatus(), info.GetUpdateTime(), info.GetId())
	if err != nil {
		common.Errorf("fail to exec sql: %s, error: %v", sql, err)
		code = common.ERR_DATABASE
		return
	}

	return
}
