package clients

import (
	"dubbo-demo/pkg/dto"
	"dubbo-demo/pkg/service"
	"dubbo.apache.org/dubbo-go/v3/config"
	hessian "github.com/apache/dubbo-go-hessian2"
)

var(
	UsZk1 = &UserServiceEnv1{}
	UsZk2 = &UserServiceEnv2{}
)

func InitUserServiceClient(){
	dto.RegisterDtos(func(pojos ...dto.POJO){
		for _,p := range pojos{
			hessian.RegisterPOJO(p)
		}
	})
	config.SetConsumerService(UsZk1)
	config.SetConsumerService(UsZk2)
}

type UserServiceEnv1 struct {
	service.UserService
}

type UserServiceEnv2 struct {
	service.UserService
}

func(a *UserServiceEnv1)Reference()string{
	return "UserServiceEnv1"
}

func(a *UserServiceEnv2)Reference()string{
	return "UserServiceEnv2"
}
