package main

import (
	"flag"
)

var (
	addr = flag.String("addr","9005","server address")
	etcdAddr = flag.String("etcdAddr","localhost:2379","etcd address")
	basePath = flag.String("basePath","/go-admin/systemService","etcd basePath")
	//servicePath = flag.String("servicePath","")
)
func main()  {
	flag.Parse()
	//handler.Init(basePath,)
}
