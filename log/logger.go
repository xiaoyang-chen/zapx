package log

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _logger *zap.Logger

type StInitFuncParams struct {
	ServiceName, FileLogLevel, ConsoleLogLevel string
	LumberjackLogger                           *lumberjack.Logger
	IsConsoleEnable                            bool
}

func Init(params StInitFuncParams) {

	// check params
	if params.LumberjackLogger == nil && !params.IsConsoleEnable {
		_logger = zap.NewNop()
		return
	}
	// 设置日志级别, "debug","info","warn","error","dPanic","panic","fatal", default "info"
	var logLevelMap = map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel, // 0
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dPanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}
	var encoderConfig = zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "line_num",
		FunctionKey:    "func",
		StacktraceKey:  "stack_trace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	// make zapcore
	var cores = make([]zapcore.Core, 0, 2)
	var fileLogLevel, consoleLogLevel = logLevelMap[params.FileLogLevel], logLevelMap[params.ConsoleLogLevel]
	if params.LumberjackLogger != nil {
		cores = append(cores, zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(params.LumberjackLogger), // 日志分割
			zap.LevelEnablerFunc(func(lvl zapcore.Level) bool { return lvl >= fileLogLevel }),
		))
	}
	if params.IsConsoleEnable {
		cores = append(cores, zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.Lock(os.Stdout),
			zap.LevelEnablerFunc(func(lvl zapcore.Level) bool { return lvl >= consoleLogLevel }),
		))
	}
	// 开启文件及行号, 跳过封装的日志函数, 设置初始化字段, 添加服务名称
	_logger = zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.AddCallerSkip(1), zap.Fields(zap.String("service_name", params.ServiceName)))
	Info("logger init success")
	// var testErr = func() error { return errors.New("testErr") } // errors == "github.com/pkg/errors"
	// Info("test", String("string", "aaa"), Int("int", 1), Error2Field(testErr()))
	// log file will show like below
	// {"level":"INFO","time":"2022-05-09T09:08:50.867+0800","line_num":"xxx/logger.go:73","func":"xxx/logger.Info","msg":"test","service_name":"xxx","string":"aaa","int":1,"error":"testErr","errorVerbose":"testErr\nxxx/chenxiaoyang/xxx.func1\n\t/xxx/logger.go:61\n/xxx/logger.Init\n\t/xxx/logger.go:62\nmain.main.func1\n\t/xxx/main.go:14\nruntime.goexit\n\t/xxx/go/1.18.1/libexec/src/runtime/asm_amd64.s:1571"}
}

func Debug(msg string, fields ...Field)  { _logger.Debug(msg, fields...) }
func Info(msg string, fields ...Field)   { _logger.Info(msg, fields...) }
func Warn(msg string, fields ...Field)   { _logger.Warn(msg, fields...) }
func Error(msg string, fields ...Field)  { _logger.Error(msg, fields...) }
func DPanic(msg string, fields ...Field) { _logger.DPanic(msg, fields...) }
func Panic(msg string, fields ...Field)  { _logger.Panic(msg, fields...) }
func Fatal(msg string, fields ...Field)  { _logger.Fatal(msg, fields...) }
