package custom

import (
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/serverplugin"
	"log"
	"time"
)

func EtcdRegistryPlugin(addr *string, etcdAddr *string, prefix *string) *serverplugin.EtcdRegisterPlugin {

	r := &serverplugin.EtcdRegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		EtcdServers:    []string{*etcdAddr},
		BasePath:       *prefix,
		Metrics:        metrics.NewRegistry(),
		Services:       nil,
		UpdateInterval: time.Minute,
		Options:        nil,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	return r
}
