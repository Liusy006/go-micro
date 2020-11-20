package main

import (
	"fmt"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"time"
)

func main(){
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))

	for{
		services, err := etcdReg.GetService("productservice")
		if err != nil{
			fmt.Println(err)
			return
		}
		next := selector.Random(services)
		node, err := next()
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Println(node.Address, " : ", node.Id, " : ", node.Metadata)
		time.Sleep(time.Second)
	}

}
