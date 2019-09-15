package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"github.com/smallnest/rpcx/client"
	"go-admin/common/middleware"
	"go-admin/system/auth/auth-api/handler"
)

var (
	addr        = flag.String("addr", ":9004", "api address")
	prefixPath  = flag.String("prefixPath", "/go-admin/auth", "etcd prefix path")
	etcdAddr    = flag.String("etcdAddr", "localhost:2379", "etcd address")
	servicePath = flag.String("serverPath", "authService", "service path, default service name")
)

func main() {
	c := client.NewEtcdDiscovery(*prefixPath, *servicePath, []string{*etcdAddr}, nil)
	xc := client.NewXClient(*servicePath, client.Failtry, client.RandomSelect, c, client.DefaultOption)
	defer xc.Close()
	cc := client.NewEtcdDiscovery("/go-admin/verification-code", "codeService", []string{*etcdAddr}, nil)
	cxc := client.NewXClient(*servicePath, client.Failtry, client.RandomSelect, cc, client.DefaultOption)
	defer cxc.Close()
	handler.SetXc(xc)
	handler.SetCodeCli(cxc)
	r := gin.Default()
	r.Use(middleware.Cors())
	r.POST("/login", handler.Login)
	r.POST("/checkToken", handler.CheckToken)
	err := r.Run(*addr)
	if err != nil {
		log.Fatalf("run server start error %v", err)
	}
}
