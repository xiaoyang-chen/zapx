package log

import (
	"sync"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// "debug","info","warn","error","dPanic","panic","fatal", default "info"

var Debug, Info, Warn, Error, DPanic, Panic, Fatal func(msg string, fields ...Field) // these value can't not be Change
var _mu sync.Mutex                                                                   // _mu mutex Debug, Info, Warn, Error, DPanic, Panic, Fatal, _logger
var _logger *zap.Logger

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
	setLoggerAndLogFunc(zap.New(core, caller, filed))
	Info("logger init success")
	// var testErr = func() error { return errors.New("testErr") } // errors == "github.com/pkg/errors"
	// Info("test", String("string", "aaa"), Int("int", 1), Error2Field(testErr()))
	// log file will show like below
	// {"level":"INFO","time":"2022-05-09T09:08:50.867+0800","line_num":"xxx/logger.go:73","func":"xxx/logger.Info","msg":"test","service_name":"xxx","string":"aaa","int":1,"error":"testErr","errorVerbose":"testErr\nxxx/chenxiaoyang/xxx.func1\n\t/xxx/logger.go:61\n/xxx/logger.Init\n\t/xxx/logger.go:62\nmain.main.func1\n\t/xxx/main.go:14\nruntime.goexit\n\t/xxx/go/1.18.1/libexec/src/runtime/asm_amd64.s:1571"}
}

func setLoggerAndLogFunc(logger *zap.Logger) {

	_mu.Lock()
	defer _mu.Unlock()

	_logger = logger
	Debug = _logger.Debug
	Info = _logger.Info
	Warn = _logger.Warn
	Error = _logger.Error
	DPanic = _logger.DPanic
	Panic = _logger.Panic
	Fatal = _logger.Fatal
}

// 包装函数会导致打印出来的调用函数显示为以下包装函数的名称
// func Debug(msg string, fields ...Field) { _logger.Debug(msg, fields...) }
// func Info(msg string, fields ...Field)   { _logger.Info(msg, fields...) }
// func Warn(msg string, fields ...Field)   { _logger.Warn(msg, fields...) }
// func Error(msg string, fields ...Field)  { _logger.Error(msg, fields...) }
// func DPanic(msg string, fields ...Field) { _logger.DPanic(msg, fields...) }
// func Panic(msg string, fields ...Field)  { _logger.Panic(msg, fields...) }
// func Fatal(msg string, fields ...Field)  { _logger.Fatal(msg, fields...) }
