package stackError

import "encoding/json"

type StatusCode struct {
	code int
	msg  string
}

func (s *StatusCode) Code() int {
	return s.code
}

func (s *StatusCode) Msg() string {
	return s.msg
}

func (s *StatusCode) String() string {
	m := map[string]interface{}{
		"code": s.Code(),
		"msg":  s.Msg(),
	}
	bytes, _ := json.Marshal(m)
	return string(bytes)
}

// 101 标识数据库错误
var (
	StatusDBGeneral      = StatusCode{10100, "服务器开小差啦"}
	StatusDBDuplicateKey = StatusCode{10101, "重复主键"}
	StatusDBInvalidSQL   = StatusCode{10102, "SQL语句非法"}
)
