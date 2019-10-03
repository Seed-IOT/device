package constant

import (
	"encoding/json"
	"fmt"
)

// Error x
type Error struct {
	Code    int
	Message string
}

var SUCCESS = 0
var UNKNOWN_ERROR = 8000
var DEVICE_NOT_FOUND = 7001

// ErrorCodes x
var ErrorCodes = map[int]string{
	SUCCESS:          "成功",
	UNKNOWN_ERROR:    "未知错误",
	DEVICE_NOT_FOUND: "设备不存在",
}

func (e Error) String() string {
	return ErrorCodes[e.Code]
}

// UnmarshalJSON x
func (e *Error) UnmarshalJSON() (Error, error) {
	errData := Error{}
	err := json.Unmarshal([]byte(fmt.Sprintf(`{"code": %d, "message": "%s"}`, e.Code, e.String())), &errData)
	return errData, err
}
