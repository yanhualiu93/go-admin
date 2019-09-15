package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/smallnest/rpcx/client"
	"go-admin/common/middleware"
	"go-admin/verification-code/code-api/handler"
)

var (
	etcdAddr = flag.String("etcdAddr", "localhost:2379", "etcd address")
	basePath = flag.String("base", "/go-admin/verification-code", "prefix path")
)

func main() {
	flag.Parse()
	c := client.NewEtcdDiscovery(*basePath, "codeService", []string{*etcdAddr}, nil)
	xclient := client.NewXClient("codeService", client.Failtry, client.RandomSelect, c, client.DefaultOption)
	handler.SetCli(xclient)
	defer xclient.Close()
	router := gin.Default()
	router.Use(middleware.Cors())
	router.GET("/verificationCode", handler.CreateCode)
	router.POST("/checkCode", handler.CheckCode)
	router.Run(":9000")
}
