package main

import (
	"dubbo-demo/cmd/server/service"
	"dubbo-demo/pkg/dto"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	hessian "github.com/apache/dubbo-go-hessian2"
)

func main(){
	config.SetProviderService(&service.UserService{})
	dto.RegisterDtos(func(pojos ...dto.POJO){
		for _,p := range pojos{
			hessian.RegisterPOJO(p)
		}
	})
	err := config.Load(config.WithPath("./conf/dubbo.yml"))
	if err != nil{
		panic(err)
	}
	select {
	}
}
