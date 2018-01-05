package log

import (
	"os"
	"path"

	"code.google.com/p/log4go"
)

var (
	clevel     = log4go.Level(LoggerLevel)
	cpath      = createLogDir(LoggerFilePath)
	cformatc   = LoggerConsoleFormat
	cformatf   = LoggerFileFormat
	csize      = LoggerFileMaxSize
	clines     = LoggerFileMaxLines
	cmaxbackup = LoggerFileBackup
)

func createLogDir(dir string) string {
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			if err = os.Mkdir(dir, os.ModePerm); err != nil {
				return "."
			} else {
				return dir
			}
		} else {
			// other error
			return "."
		}
	} else {
		//exist
		return dir
	}
}

var (
	wfroot  = newFileLogWriter("root")
	wferror = newFileLogWriter("error")
	wc      = newConsoleLogWriter()
)

var (
	// Root Logger
	Root logger = log4go.Logger{
		"froot":  &log4go.Filter{clevel, wfroot},
		"ferror": &log4go.Filter{log4go.ERROR, wferror},
	}
)

// Quick use for Root.(method)
var (
	Finest   = Root.Finest
	Fine     = Root.Fine
	Debug    = Root.Debug
	Trace    = Root.Trace
	Info     = Root.Info
	Warn     = Root.Warn
	Error    = Root.Error
	Critical = Root.Critical
)

type logger interface {
	Finest(arg0 interface{}, args ...interface{})
	Fine(arg0 interface{}, args ...interface{})
	Debug(arg0 interface{}, args ...interface{})
	Trace(arg0 interface{}, args ...interface{})
	Info(arg0 interface{}, args ...interface{})
	Warn(arg0 interface{}, args ...interface{}) error
	Error(arg0 interface{}, args ...interface{}) error
	Critical(arg0 interface{}, args ...interface{}) error
}

func newFileLogWriter(name string) *log4go.FileLogWriter {
	fname := path.Join(cpath, name+".log")
	flw := log4go.NewFileLogWriter(fname, false)
	flw.SetFormat(cformatf)
	flw.SetRotate(false)
	flw.SetRotateSize(csize)
	flw.SetRotateLines(clines)
	flw.SetRotateDaily(true)
	flw.SetRotateMaxBackup(cmaxbackup)
	return flw
}

func newConsoleLogWriter() *log4go.ConsoleLogWriter {
	clw := log4go.NewConsoleLogWriter()
	clw.SetFormat(cformatc)
	return clw
}

func init() {
	debugMode := LoggerModeDebug
	if debugMode {
		// Root Logger
		Root = log4go.Logger{
			"stdout": &log4go.Filter{clevel, wc},
			"froot":  &log4go.Filter{clevel, wfroot},
			"ferror": &log4go.Filter{log4go.ERROR, wferror},
		}
	}
}
