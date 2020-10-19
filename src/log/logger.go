/**
 * Created by 2020/10/20.
 */

package log

import (
	"github.com/alecthomas/log4go"
)

//<!-- level is (:?FINEST|FINE|DEBUG|TRACE|INFO|WARNING|ERROR) -->
var logger = log4go.NewLogger()

func Init() {
	logger.LoadConfiguration("./log.xml")

	//test
	//LOGGER.Finest("Finest %v", time.Now())
	//LOGGER.Fine("Fine %v", time.Now())
	//LOGGER.Debug("Debug %v", time.Now())
	//LOGGER.Trace("Trace %v", time.Now())
	//LOGGER.Info("Info %v", time.Now())
	//LOGGER.Warn("Warn %v", time.Now())
	//LOGGER.Error("Error %v", time.Now())
}

func Close() {
	logger.Close()
}

//做一层适配，方便后续切换到其他日志框架或者自己写

func Finest(arg0 interface{}, args ...interface{}) {
	logger.Finest(arg0, args...)
}

func Fine(arg0 interface{}, args ...interface{}) {
	logger.Fine(arg0, args...)
}

func Debug(arg0 interface{}, args ...interface{}) {
	logger.Debug(arg0, args...)
}

func Trace(arg0 interface{}, args ...interface{}) {
	logger.Trace(arg0, args...)
}

func Info(arg0 interface{}, args ...interface{}) {
	logger.Info(arg0, args...)
}

func Warn(arg0 interface{}, args ...interface{}) {
	logger.Warn(arg0, args...)
}

func Error(arg0 interface{}, args ...interface{}) {
	logger.Error(arg0, args...)
}

func Critical(arg0 interface{}, args ...interface{}) {
	logger.Critical(arg0, args...)
}
