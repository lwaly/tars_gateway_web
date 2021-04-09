package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lwaly/tars_gateway_web/server/common"
	"github.com/lwaly/tars_gateway_web/server/routers"
)

func main() {
	var strConfPath string

	flag.StringVar(&strConfPath, "c", "", "The server config path")
	if err := common.NewConf(strConfPath); nil != err {
		fmt.Println(err)
		fmt.Println(`
		Usage: gateway [options]
		
		Options:
			-c, <config path>
		`)
		os.Exit(1)
	}

	routers.Route()
}
