package log

// 追加标签
func (l *Logger) AppendMark(mark string) {
	l.mark += mark
}

// 设置标签
func (l *Logger) SetMark(mark string) {
	l.mark = mark
}

// 设置标签
func SetMark(mark string) {
	GetLogout().SetMark(mark)
}

// 获取当前标签
func (l *Logger) GetMark() string {
	return l.mark
}

// 标签
func WithMark(mark string) LoggerOption {
	return func(l *Logger) {
		l.SetMark(mark)
	}
}

// 关闭标签
func WithDisableMark() LoggerOption {
	return func(l *Logger) {
		l.disableMark = true
	}
}

func DisableMark() {
	GetLogout().DisableMark()
}

func (l *Logger) DisableMark() {
	l.disableMark = true
}

func SetShortout(value bool) {
	GetLogout().SetShortout(value)
}

func WithShortout() LoggerOption {
	return func(l *Logger) {
		l.short = true
	}
}

func Shortout() {
	GetLogout().Shortout()
}

func (l *Logger) Shortout() {
	l.short = true
}

func (l *Logger) SetShortout(value bool) {
	l.short = value
}
