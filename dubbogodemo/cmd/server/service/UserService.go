package service

import (
	"context"
	"dubbo-demo/pkg/dto"
	util2 "dubbo-demo/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"time"
)

// UserService
//  UserService
//  @Description: dubbo service 的具体实现
//
type UserService struct {
}

func (u *UserService) Reference() string {
	return "UserService"
}

func (u *UserService) QueryUser(ctx context.Context, in *dto.UserRequest)(*dto.UserResponse, error){
	appName, traceId := util2.GetDubboContextAppName(ctx)
	logrus.WithFields(logrus.Fields{
		"appName" : appName,
		"traceId": traceId,
		"req": in,
	}).Info("user_service:queryuser received call")
	if appName == ""{
		return nil, errors.Errorf("auth error, no appname")
	}
	if appName != "dubbo-consumer"{
		return nil, errors.Errorf("auth error, not allowed")
	}
	if in.UserId != ""{
		if in.UserId != "123456789"{
			return nil, errors.Errorf("no user found with id[%s]", in.UserId)
		}
		return &dto.UserResponse{
			UserName:  "dubbo-user1",
			UserId:    "123456789",
			BirthDate: util2.GetCommonTimeStr(time.Now().Add(-56 * 360 * 24 * time.Hour)),
			Age:       56,
			IsDead:    true,
		}, nil
	}
	if in.UserName == ""{
		return nil, errors.Errorf("please input query condition")
	}
	return &dto.UserResponse{
		UserName:  in.UserName,
		UserId:    "666666666",
		BirthDate: util2.GetCommonTimeStr(time.Now().Add(-30 * 360 * 24 * time.Hour)),
		Age:       30,
		IsDead:    false,
	}, nil
}
