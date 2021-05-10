package rotate

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

// 每隔1天生成一个新的日志文件，保留7天，文件最大为1024*1024*10
func NewWriter(logPath string) (*rotatelogs.RotateLogs, error) {
	return rotatelogs.New(
		logPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithMaxAge(time.Duration(60*60*24)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(60*60*24*7)*time.Second),
		rotatelogs.WithRotationSize(1024*1024*10),
	)
}
