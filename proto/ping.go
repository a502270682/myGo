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
