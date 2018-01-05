package log

const (
	// LoggerModeDebug : debug mode
	LoggerModeDebug bool = true
	// LoggerLevel : FINEST ,FINE ,DEBUG ,TRACE ,INFO ,WARNING, ERROR, CRITICAL -> 0 ~ 7
	LoggerLevel int = 2
	// LoggerConsoleFormat : logger console print format
	LoggerConsoleFormat string = "[%T %D] [%L] (%S) %M"
	// LoggerFileFormat : logger file print format
	LoggerFileFormat string = "[%T %D] [%L] (%S) %M"
	// LoggerFilePath : logger file path
	LoggerFilePath string = "./log"
	// LoggerFileMaxSize : logger file max size
	LoggerFileMaxSize int = 1048576
	// LoggerFileMaxLines : logger file max lines
	LoggerFileMaxLines int = 10000
	// LoggerFileBackup : logger file backup number
	LoggerFileBackup int = 10
)
