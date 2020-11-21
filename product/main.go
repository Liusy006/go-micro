package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-micro/registry/etcd"
)

func main(){
	ginRouter := gin.Default()
	v1Group := ginRouter.Group("v1")
	{
		v1Group.Handle("POST", "/product", func(context *gin.Context) {
			context.JSON(200,
				gin.H{
					"data": newProductList(5),
				},
				)
		})
	}


	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))

	service := web.NewService(
		web.Name("productservice"),
		//web.Address(":8888"),
		web.Handler(ginRouter),
		web.Registry(etcdReg),
		)

	service.Init()
	service.Run()

}
