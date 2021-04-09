package mymysql

import (
	"time"

	"github.com/lwaly/tars_gateway_web/server/common"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var mgo_idle_count = 200   //连接池空闲个数
var mgo_active_count = 250 //连接池活动个数
var mgo_idle_timeout = 10  //空闲超时时间
var externalIp string

type StMysqlConf struct {
	Addr           string `json:"addr,omitempty"`
	MgoIdleCount   int    `json:"mgo_idle_count,omitempty"`   //最大空闲连接数
	MgoActivecount int    `json:"mgo_active_count,omitempty"` //最大活动连接数
	MgoIdleTimeout int    `json:"mgo_idle_timeout,omitempty"` //连接超时时间
}

var mysqlClients map[int]*sql.DB
var oMysqlClients map[string]map[string]*sql.DB

func InitMysql() {
	mysqlClients = make(map[int]*sql.DB)
	var stMysqlConf StMysqlConf
	err := common.Conf.GetStruct("http", &stMysqlConf)
	if err != nil {
		common.Errorf("fail to get http conf.%v", err)
		return
	}

	common.Infof("%v", stMysqlConf)

	index := 0

	if obj := createPool(stMysqlConf.MgoIdleCount, stMysqlConf.MgoActivecount, stMysqlConf.MgoIdleTimeout, stMysqlConf.Addr); nil != obj {
		mysqlClients[index] = obj
		index++
	}

}

func createPool(maxIdle, maxActive, idleTimeout int, address string) *sql.DB {
	//address = fmt.Sprintf("root:root@appinside@tcp(192.168.0.20:3306)/db_tars?charset=utf8")
	obj, err := sql.Open("mysql", address)
	common.Infof("%s", address)
	if nil != err {
		common.Errorf("%v", err)
		return nil
	}
	obj.SetConnMaxLifetime(time.Duration(idleTimeout))
	obj.SetMaxIdleConns(maxIdle)
	obj.SetMaxOpenConns(maxActive)
	//obj.Exec("SET NAMES 'utf8mb4'; SET CHARACTER SET utf8mb4;")
	return obj
}

func GetConn() *sql.DB {
	return mysqlClients[0]
}
