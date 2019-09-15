package main

import (
	"flag"
	"github.com/go-redis/redis"
	"github.com/smallnest/rpcx/server"
	"go-admin/common/custom"
	"go-admin/verification-code/code-server/model"
	"go-admin/verification-code/code-server/service"
)

var (
	addr      = flag.String("addr", "localhost:8973", "server address")
	etcdAddr  = flag.String("etcdAddr", "localhost:2379", "etcd address")
	basePath  = flag.String("base", "/go-admin/verification-code", "prefix path")
	redisAddr = flag.String("redisAddr", "47.104.185.237:6379", "redis address")
	redisPass = flag.String("redisPass", "17185383573", "redis password")
)

func main() {
	flag.Parse()
	model.Init(redis.NewClient(&redis.Options{
		Addr:     *redisAddr,
		Password: *redisPass,
		DB:       0,
	}), "", 0)
	s := server.NewServer()
	s.Plugins.Add(custom.EtcdRegistryPlugin(addr, etcdAddr, basePath))
	s.RegisterName("codeService", new(service.Code), "")
	//s.Register(new(service.Code), "")
	s.Serve("tcp", *addr)
}
