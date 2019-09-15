package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	"go-admin/common/custom"
	"go-admin/system/system-server/service"
)

var(
	addr = flag.String("addr","8821","server address")
	etcdAddr =flag.String("etcdAddr","localhost:2379","etcd address")
	prefix  = flag.String("prefix","/go-admin/system-service","etcd prefix path")
)

func main()  {
	s := server.NewServer()
	s.Plugins.Add(custom.EtcdRegistryPlugin(addr,etcdAddr,prefix))
	_=s.RegisterName("systemService",&service.SysService{},"")
	_=s.Serve("tcp",*addr)
}
