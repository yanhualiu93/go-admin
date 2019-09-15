package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	"github.com/xormplus/core"
	"go-admin/common/custom"
	"go-admin/system/auth/auth-server/service"
	"go-admin/system/models"
)

var (
	addr       = flag.String("addr", "localhost:8888", "server address")
	etcdAddr   = flag.String("etcdAddr", "localhost:2379", "etcd address")
	etcdPrefix = flag.String("etcdPrefix", "/go-admin/auth", "etcd prefix path")
)

func main() {
	flag.Parse()
	e := custom.NewDb("root", "123456", "47.104.185.237", "3306", "go-admin")
	e.SetLogLevel(core.LOG_DEBUG)
	models.SetDB(e)
	ser := server.NewServer()
	ser.Plugins.Add(custom.EtcdRegistryPlugin(addr, etcdAddr, etcdPrefix))
	ser.RegisterName("authService", &service.TokenService{}, "")
	ser.Serve("tcp", *addr)
}
