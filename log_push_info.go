package log

import "github.com/bytedance/sonic"

// 日志信息
type LogInfo struct {
	CreateTime int64  `json:"CreateTime"` // 日志创建时间
	CodeFile   string `json:"CodeFile"`   // 代码文件
	CodeLine   int    `json:"CodeLine"`   // 代码行
	Level      int    `json:"Level"`      // 日志等级
	Content    string `json:"Content"`    // 日志内容
}

// 转为JSON字符串
func (l *LogInfo) ToJSONString() string {
	res, _ := sonic.MarshalString(l)
	return res
}

func (l *LogInfo) ToJSONByte() []byte {
	res, _ := sonic.Marshal(l)
	return res
}
