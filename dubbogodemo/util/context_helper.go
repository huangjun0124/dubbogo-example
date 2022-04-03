package util

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
)

func GetDubboContextAppName(ctx context.Context)(appName string){
	ctxMap, ok := ctx.Value(constant.DubboCtxKey("attachment")).(map[string]interface{})
	if !ok{
		return
	}
	if str, ok := ctxMap["app"]; ok{
		appName, _ = str.(string)
	}
	return
}
