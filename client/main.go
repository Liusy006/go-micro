package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"io/ioutil"
	"log"
	"net/http"
	my_http "github.com/micro/go-plugins/client/http"
	"time"
)

func callApi2( s selector.Selector)  {
	myclient := my_http.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
		)

	req:= myclient.NewRequest("productservice", "/v1/product", nil)

	var rsp map[string]interface{}
	err := myclient.Call(context.Background(), req,  &rsp)
	if err != nil{
		log.Fatal("****", err)
	}

	fmt.Println(rsp["data"])
}

func callApi( addr, path, method string, body []byte)(string, error){
	req, err := http.NewRequest(method, "http://"+addr + path, bytes.NewReader(body))
	if err != nil{
		log.Fatal(err)
	}
	client := http.DefaultClient;
	reps, err := client.Do(req)
	if err != nil{
		log.Fatal(err)
	}

	defer reps.Body.Close()
	d, err := ioutil.ReadAll(reps.Body)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(string(d))
	return string(d), nil
}

func main(){
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	s := selector.NewSelector(
		selector.Registry(etcdReg),
		selector.SetStrategy(selector.RoundRobin),
		)

	callApi2( s)
	time.Sleep(time.Second)

	/*for{
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
		callApi(node.Address, "/v1/product", "POST", nil)
		time.Sleep(time.Second)
	}*/

}
