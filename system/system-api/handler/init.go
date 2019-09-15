package handler

import "github.com/smallnest/rpcx/client"

var Xclient client.XClient
func Init(basePath string,servicePath string,etcdAddr []string)  {
	c := client.NewEtcdDiscovery(basePath,servicePath,etcdAddr,nil)
	Xclient=client.NewXClient(servicePath,client.Failtry,client.RandomSelect,c,client.DefaultOption)
}
