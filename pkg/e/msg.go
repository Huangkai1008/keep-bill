package e

// 错误码与信息对照
var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "请求参数错误",
	ErrorDataBase: "数据库错误",
	RepeatUser:    "重复的用户",
	NotExistUser:  "不存在的用户",
	ErrorAuth:     "登录信息有误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[Error]
}
