package proto

type CommonRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PingReq struct {
}

type PingRsp struct {
	Success string `json:"success"`
}

type UserInfoReq struct {
	Name string `form:"name"`
}

type UserInfoRsp struct {
	UserName string `json:"user_name"`
	UserAge  int    `json:"user_age"`
}
