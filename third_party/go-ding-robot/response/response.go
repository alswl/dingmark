package response

type Response struct {
	Code    int    `json:"errcode"`
	Message string `json:"errmsg"`
	Status  int    `json:"status"`
	Punish  string `json:"punish"`
	Wait    int    `json:"wait"`
}

func (r Response) IsSuccess() bool {
	return r.GetCode() == 0
}

func (r Response) GetCode() int {
	return r.Code
}

func (r Response) GetMessage() string {
	return r.Message
}
