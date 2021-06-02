package handlers

import (
	"context"
	"myGo/adapter/error_code"
	"myGo/proto"
)

func SchoolListHandler(ctx context.Context, req *proto.SchoolListReq, rsp *proto.SchoolListRsp) *error_code.ReplyError {
	return nil
}
