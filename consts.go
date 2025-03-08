package log

type (
	LogLevel  = int
	LogPrefix = string
)

const (
	// 日志默认前缀
	LogDefaultPrefix = "DEFAULT"

	// 日志等级前缀
	LogPrefixColorSpeed LogPrefix = "\033[01;36mSPEED\033[00m "
	LogPrefixColorDebug LogPrefix = "\033[01;32mDEBUG\033[00m "
	LogPrefixColorInfo  LogPrefix = "\033[01;34mINFO\033[00m  "
	LogPrefixColorWarn  LogPrefix = "\033[01;33mWARN\033[00m  "
	LogPrefixColorError LogPrefix = "\033[01;31mERROR\033[00m "

	LogPrefixNoColorSpeed LogPrefix = "SPEED "
	LogPrefixNoColorDebug LogPrefix = "DEBUG "
	LogPrefixNoColorInfo  LogPrefix = "INFO  "
	LogPrefixNoColorWarn  LogPrefix = "WARN  "
	LogPrefixNoColorError LogPrefix = "ERROR "

	// 日志等级
	SPEEDLevel LogLevel = 5
	DEBUGLevel LogLevel = 4
	INFOLevel  LogLevel = 3
	WARNLevel  LogLevel = 2
	ERRORLevel LogLevel = 1

	LogLevelSpeed LogLevel = 5
	LogLevelDebug LogLevel = 4
	LogLevelInfo  LogLevel = 3
	LogLevelWarn  LogLevel = 2
	LogLevelError LogLevel = 1
)

var LogLevelColorPrefixMap = map[LogLevel]LogPrefix{
	LogLevelSpeed: LogPrefixColorSpeed,
	LogLevelDebug: LogPrefixColorDebug,
	LogLevelInfo:  LogPrefixColorInfo,
	LogLevelWarn:  LogPrefixColorWarn,
	LogLevelError: LogPrefixColorError,
}

var LogLevelNoColorPrefixMap = map[LogLevel]LogPrefix{
	LogLevelSpeed: LogPrefixNoColorSpeed,
	LogLevelDebug: LogPrefixNoColorDebug,
	LogLevelInfo:  LogPrefixNoColorInfo,
	LogLevelWarn:  LogPrefixNoColorWarn,
	LogLevelError: LogPrefixNoColorError,
}

var LogerLevel int
