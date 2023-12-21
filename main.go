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

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	svcCtx := svc.GetInitSvcContext()

	handler.RegisterHandler(server, &svcCtx)
	server.Start()

	fmt.Println(`running server at port 9101`)

}
