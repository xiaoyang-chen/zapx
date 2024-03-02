package log

import (
	"testing"

	"github.com/natefinch/lumberjack"
)

func TestInit(t *testing.T) {
	type args struct {
		params StInitFuncParams
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "only file log",
			args: args{
				params: StInitFuncParams{
					ServiceName:  "test-log",
					FileLogLevel: "info",
					// ConsoleLogLevel: "info",
					LumberjackLogger: &lumberjack.Logger{
						Filename:   "./log/test-log.log", // 日志文件路径，默认 os.TempDir()
						MaxSize:    10,                   // 每个日志文件保存10M，默认 100M
						MaxAge:     30,                   // 保留30天，默认不限
						MaxBackups: 30,                   // 保留30个备份，默认不限
						Compress:   true,                 // 是否压缩，默认不压缩
					},
					IsConsoleEnable: false,
				},
			},
		},
		{
			name: "only console log",
			args: args{
				params: StInitFuncParams{
					ServiceName: "test-log",
					// FileLogLevel: "info",
					ConsoleLogLevel:  "info",
					LumberjackLogger: nil,
					IsConsoleEnable:  true,
				},
			},
		},
		{
			name: "file and console log",
			args: args{
				params: StInitFuncParams{
					ServiceName:     "test-log",
					FileLogLevel:    "info",
					ConsoleLogLevel: "info",
					LumberjackLogger: &lumberjack.Logger{
						Filename:   "./log/test-log.log", // 日志文件路径，默认 os.TempDir()
						MaxSize:    10,                   // 每个日志文件保存10M，默认 100M
						MaxAge:     30,                   // 保留30天，默认不限
						MaxBackups: 30,                   // 保留30个备份，默认不限
						Compress:   true,                 // 是否压缩，默认不压缩
					},
					IsConsoleEnable: true,
				},
			},
		},
		{
			name: "no op log",
			args: args{
				params: StInitFuncParams{
					ServiceName:      "",
					FileLogLevel:     "",
					ConsoleLogLevel:  "",
					LumberjackLogger: nil,
					IsConsoleEnable:  false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init(tt.args.params)
		})
	}
}
