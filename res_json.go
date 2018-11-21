package common

type ResJson struct {
	Success bool `json:"success"`
	Code    int  `json:"code"`
	Msg     *string `json:"msg"`
	Data    interface{} `json:"data"`
}

func BuildSuccess(data interface{}) ResJson{
	return ResJson{Success:true,Code:0,Msg:nil,Data:data}
}

func BuildFailure(code int,msg string,data interface{}) ResJson{
	return ResJson{Success:false,Code:code,Msg:&msg,Data:data}
}
