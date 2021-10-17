package cmd

import (
	"dubbo-demo/pkg/dto"
	"dubbo-demo/pkg/service"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"

	hessian "github.com/apache/dubbo-go-hessian2"
)

func StartDubboServer(){
	config.SetProviderService(&service.UserService{})
	hessian.RegisterPOJO(&dto.UserRequest{})
	hessian.RegisterPOJO(&dto.UserResponse{})

	err := config.Load(config.WithPath("./conf/dubbo.yml"))
	if err != nil{
		panic(err)
	}
	select {
	}
}
