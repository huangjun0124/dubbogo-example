package main

import (
	"dubbo-demo/cmd/client/clients"
	"dubbo-demo/pkg/dto"
	"dubbo-demo/util"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/sirupsen/logrus"
)

func main(){
	clients.InitUserServiceClient()
	err := config.Load(config.WithPath("./conf/dubbo.yml"))
	if err != nil{
		panic(err)
	}
	req := &dto.UserRequest{
		UserName: "你好",
		UserId:   "123",
	}
	ctx, traceId := util.GetDubboContextWithAppName("dubbo-consumer")
	//time.Sleep(10 * time.Second)
	rsp, err := clients.UsZk1.QueryUser(ctx, req)
	if err != nil{
		logrus.WithFields(logrus.Fields{
			"traceId": traceId,
			"rsp": rsp,
			"err": err,
		}).Error("query zk1 userservice response")
	}else{
		logrus.WithFields(logrus.Fields{
			"traceId": traceId,
			"rsp": rsp,
		}).Info("query zk1 userservice response")
	}

	req.UserId = "123456789"
	ctx, traceId = util.GetDubboContextWithAppName("dubbo-consumer")
	rsp, err = clients.UsZk2.QueryUser(ctx, req)
	if err != nil{
		logrus.WithFields(logrus.Fields{
			"traceId": traceId,
			"rsp": rsp,
			"err": err,
		}).Error("query zk2 userservice response")
	}else{
		logrus.WithFields(logrus.Fields{
			"traceId": traceId,
			"rsp": rsp,
		}).Info("query zk2 userservice response")
	}
	select {
	}
}
