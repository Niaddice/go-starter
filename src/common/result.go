package common

type Msg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type MsgWord string

const (
	IntervalError        MsgWord = "内部错误"
	RequestParamError    MsgWord = "请求参数不正确"
	JsonConvertError     MsgWord = "Json转换错误"
	ApiRequestError      MsgWord = "接口请求错误"
	ApiResponseDataError MsgWord = "接口数据错误"
)

func Ok(data interface{}) Msg {
	return Msg{
		Code: 200,
		Msg:  "ok",
		Data: data,
	}
}

func Error(data interface{}) Msg {
	return Msg{
		Code: 500,
		Msg:  "failed",
		Data: data,
	}
}

func DefaultFailed(err error) Msg {
	return Msg{
		Code: 500,
		Msg:  err.Error(),
		Data: nil,
	}
}

func Failed(msg MsgWord) Msg {
	return Msg{
		Code: 500,
		Msg:  string(msg),
		Data: nil,
	}
}

func FailedByReqParam() Msg {
	return Msg{
		Code: 500,
		Msg:  string(RequestParamError),
		Data: nil,
	}
}
