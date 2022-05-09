package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _logger *zap.Logger // consider using atomic.Value for data race

func Init(serviceName, loglevel string, lumberjackLogger lumberjack.Logger) {

	// 设置日志级别, 默认info
	var logLevelMap = map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel, // 0
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dPanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}
	var level = logLevelMap[loglevel]
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
	// 设置日志级别
	var atomicLevel = zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	var core = zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(&lumberjackLogger), // 日志分割
		atomicLevel,
	)
	// 开启文件及行号
	var caller = zap.AddCaller()
	// 设置初始化字段,如：添加一个服务器名称
	var filed = zap.Fields(zap.String("service_name", serviceName))
	// 构造日志
	_logger = zap.New(core, caller, filed)
	Info("logger init success")
	// var testErr = func() error { return errors.New("testErr") } // errors == "github.com/pkg/errors"
	// Info("test", String("string", "aaa"), Int("int", 1), Error2Field(testErr()))
	// log file will show as follow
	// {"level":"INFO","time":"2022-05-09T09:08:50.867+0800","line_num":"xxx/logger.go:73","func":"xxx/logger.Info","msg":"test","service_name":"xxx","string":"aaa","int":1,"error":"testErr","errorVerbose":"testErr\nxxx/chenxiaoyang/xxx.func1\n\t/xxx/logger.go:61\n/xxx/logger.Init\n\t/xxx/logger.go:62\nmain.main.func1\n\t/xxx/main.go:14\nruntime.goexit\n\t/xxx/go/1.18.1/libexec/src/runtime/asm_amd64.s:1571"}
}

// func getLogger() *zap.Logger { return _logger.Load().(*zap.Logger) }
func getLogger() *zap.Logger { return _logger }

// "debug","info","warn","error","dPanic","panic","fatal", default "info"
func Debug(msg string, fields ...field)  { getLogger().Debug(msg, fields...) }
func Info(msg string, fields ...field)   { getLogger().Info(msg, fields...) }
func Warn(msg string, fields ...field)   { getLogger().Warn(msg, fields...) }
func Error(msg string, fields ...field)  { getLogger().Error(msg, fields...) }
func DPanic(msg string, fields ...field) { getLogger().DPanic(msg, fields...) }
func Panic(msg string, fields ...field)  { getLogger().Panic(msg, fields...) }
func Fatal(msg string, fields ...field)  { getLogger().Fatal(msg, fields...) }
