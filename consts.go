package log

const (
	// 日志默认前缀
	LogDefaultPrefix = "DEFAULT"

	// 日志等级前缀
	SPEEDPrefix = "\033[01;36mSPEED\033[00m "
	DEBUGPrefix = "\033[01;32mDEBUG\033[00m "
	INFOPrefix  = "\033[01;34mINFO\033[00m  "
	WARNPrefix  = "\033[01;33mWARN\033[00m  "
	ERRORPrefix = "\033[01;31mERROR\033[00m "

	// 日志等级
	SPEEDLevel = 5
	DEBUGLevel = 4
	INFOLevel  = 3
	WARNLevel  = 2
	ERRORLevel = 1
)

var LogLevelIntMap = map[int]string{
	SPEEDLevel: "speed",
	DEBUGLevel: "debug",
	INFOLevel:  "info",
	WARNLevel:  "warn",
	ERRORLevel: "error",
}

var LogLevelStrMap = map[string]int{
	SPEEDPrefix: SPEEDLevel,
	DEBUGPrefix: DEBUGLevel,
	INFOPrefix:  INFOLevel,
	WARNPrefix:  WARNLevel,
	ERRORPrefix: ERRORLevel,
}

var LogerLevel int
