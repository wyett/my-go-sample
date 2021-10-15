package common

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Build info
var BRANCH = "$"
var SIGNALPROFILE = "$"
var SIGNALSTACK = "$"

const (
	// APPNAME = "chatroom"
	// AppDatabase          = APPNAME
	// APPConflictDatabase  = APPNAME + "_conflict"

	GlobalDiagnosticPath = "chatlog"
	// This is the time of golang was born to the world
	GolangSecurityTime = "2006-01-02T15:04:05Z"

	WorkGood       uint64 = 0
	GetReady       uint64 = 1
	FetchBad       uint64 = 2
	TunnelSendBad  uint64 = 4
	TunnelSyncBad  uint64 = 8
	ReplicaExecBad uint64 = 16

	MajorityWriteConcern = "majority"

	Int32max = (int64(1) << 32) - 1
)

func init() {
	// prepare global folders
	Mkdirs(GlobalDiagnosticPath /*, GlobalStoragePath*/)
}

func Mkdirs(dirs ...string) error {
	for _, dir := range dirs {
		if _, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
			if err = os.Mkdir(dir, 0777); err != nil {
				return err
			}
		}
	}
	return nil
}

/*
func RunStatusMessage(status uint64) string {
	switch status {
	case WorkGood:
		return "Good"
	case GetReady:
		return "prepare for ready"
	case FetchBad:
		return "can't fetch oplog from source MongoDB"
	case TunnelSendBad:
		return "collector send oplog to tunnel failed"
	case TunnelSyncBad:
		return "receiver fetch from tunnel failed"
	case ReplicaExecBad:
		return "receiver replica executed failed"
	default:
		return "unknown"
	}
}*/
func InitialLogger(logDir, logFile, level string, logFlush bool, verbose bool) error {
	logLevel := parseLogLevel(level)
	if verbose {
		writer := LOG.NewConsoleLogWriter()
		writer.SetFormat("[%D %T] [%L] %M")
		LOG.AddFilter("console", logLevel, writer)
	}

	if len(logDir) == 0 {
		logDir = "logs"
	}
	// check directory exists
	if _, err := os.Stat(logDir); err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(logDir, os.ModeDir|os.ModePerm); err != nil {
			return fmt.Errorf("create log.dir[%v] failed[%v]", logDir, err)
		}
	}

	if len(logFile) != 0 {
		if !logFlush {
			LOG.LogBufferLength = 32
		} else {
			LOG.LogBufferLength = 0
		}
		fileLogger := LOG.NewFileLogWriter(fmt.Sprintf("%s/%s", logDir, logFile), true)
		fileLogger.SetRotateDaily(true)
		// fileLogger.SetFormat("[%D %T] [%L] [%s] %M") // print function
		fileLogger.SetFormat("[%D %T] [%L] %M")
		fileLogger.SetRotateMaxBackup(7)
		LOG.AddFilter("file", logLevel, fileLogger)
	} else {
		return fmt.Errorf("log.file[%v] shouldn't be empty", logFile)
	}

	return nil
}

func parseLogLevel(level string) LOG.Level {
	switch strings.ToLower(level) {
	case "debug":
		return LOG.DEBUG
	case "info":
		return LOG.INFO
	case "warning":
		return LOG.WARNING
	case "error":
		return LOG.ERROR
	default:
		return LOG.DEBUG
	}
}
