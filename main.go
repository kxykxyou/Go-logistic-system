package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	"logistic/internal/config"
	"logistic/internal/handler"
	"logistic/internal/svc"
)

func main() {

	c := config.GetConfigCtx()

	c.RestConf.Port = 8888
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	svcCtx := svc.GetInitSvcContext()

	handler.RegisterHandler(server, &svcCtx)
	fmt.Println(`running server at port 8888.`)
	server.Start()

}
