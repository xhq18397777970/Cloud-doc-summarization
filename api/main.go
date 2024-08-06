package api

import (
	"encoding/json"
	"fmt"
	"todo_list/serializer"
)

// 当存在错误时，将错误信息以固定的json格式输出，序列化器
func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 40001,
			Msg:    "JSON类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}

	//比如用户名为空或密码为空
	return serializer.Response{
		Status: 40001,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}
