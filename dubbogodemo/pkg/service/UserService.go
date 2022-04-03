package service

import (
	"context"
	"dubbo-demo/pkg/dto"
)

//
//  UserService
//  @Description: 服务接口定义
//
type UserService struct {
	QueryUser func(ctx context.Context, in *dto.UserRequest)(*dto.UserResponse, error)
}
