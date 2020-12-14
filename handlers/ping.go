package handlers

import (
	"context"
	"myGo/adapter/error_code"
	"myGo/proto"
)

func PingHandler(ctx context.Context, req *proto.PingReq, rsp *proto.PingRsp) *error_code.ReplyError {
	rsp.Success = "hello"
	return nil
}
