package util

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
)

func GetDubboContextAppName(ctx context.Context)(appName, traceId string){
	ctxMap, ok := ctx.Value(constant.DubboCtxKey("attachment")).(map[string]interface{})
	if !ok{
		return
	}
	if str, ok := ctxMap["app"]; ok{
		appName, _ = str.(string)
	}
	if str, ok := ctxMap["traceId"]; ok{
		traceId, _ = str.(string)
	}
	return
}

func GetDubboContextWithAppName(appName string)(context.Context, string){
	traceId := GenUuid()[0:8]
	args := make(map[string]interface{})
	args["app"] = appName
	args["traceId"] = traceId
	ctx := context.WithValue(context.Background(), constant.DubboCtxKey("attachment"), args)
	return ctx, traceId
}