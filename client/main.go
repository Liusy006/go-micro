package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	my_http "github.com/micro/go-plugins/client/http"
	"go-micro/models"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func callApi2( service, endpoit string, s selector.Selector)  {
	myclient := my_http.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
		)


	req:= myclient.NewRequest(service,
		endpoit,
		models.ProductRequest{Size:3},
		)

	var rsp models.ProductResponse
	err := myclient.Call(context.Background(), req,  &rsp)
	if err != nil{
		log.Fatal("****", err)
	}

	fmt.Println(rsp.GetData())
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

	callApi2( "productservice", "/v1/product", s)
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
