package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/micro/go-micro/v2/logger"
	cart "github.com/wxyMicro/cart/proto/cart"
	cartApi "github.com/wxyMicro/cartApi/proto/cartApi"
	"strconv"
)

type CartApi struct {
	cart.CartService
}

// CartApi.Call 通过API向外暴露为/cartApi/findAll，接收http请求
// 即：/cartApi/call请求会调用go.micro.api.cartApi 服务的CartApi.Call方法
func (e *CartApi) FindAll(ctx context.Context, req *cartApi.Request, rsp *cartApi.Response) error {
	log.Info("接收到 /cartApi/findAll 访问请求")
	if _, ok := req.Get["user_id"]; !ok {
		return errors.New("参数异常")
	}
	userIdStr := req.Get["user_id"].Values[0]
	fmt.Println(userIdStr)
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return err
	}
	//获取购物车所有商品
	cartAll, err := e.CartService.GetAll(context.TODO(), &cart.CartFindAll{UserId: userId})
	//数据类型转化
	b, err := json.Marshal(cartAll)
	if err != nil {
		return err
	}
	rsp.StatusCode = 200
	rsp.Body = string(b)
	return nil
}
