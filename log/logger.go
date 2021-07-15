package log

import (
	"bufio"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

const (
	logFileName = "lc.log"
)

func InitLogger(dir, logLevel string) error {
	logDir := filepath.Join(dir, "logs")
	_, err := os.Stat(logDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(logDir, 0755)
		if err != nil {
			return err
		}
	}

	baseLogPath := filepath.Join(logDir, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath),      // generate soft-link，point the latest log
		rotatelogs.WithMaxAge(7*24*time.Hour),     // maximum file storage time
		rotatelogs.WithRotationTime(24*time.Hour), // log cutting interval
	)
	if err != nil {
		return errors.Wrap(err, "create rotatelogs failed")
	}

	//log.SetFormatter(&log.TextFormatter{})
	switch logLevel {
	/*
	   if the log level is not debug, do not print the log to the console
	*/
	case "debug":
		log.SetLevel(log.DebugLevel)
		log.SetOutput(os.Stderr)
	case "info":
		setNull()
		log.SetLevel(log.InfoLevel)
	case "warn":
		setNull()
		log.SetLevel(log.WarnLevel)
	case "error":
		setNull()
		log.SetLevel(log.ErrorLevel)
	default:
		setNull()
		log.SetLevel(log.InfoLevel)
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer, // set different output purposes for different levels
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{})
	log.AddHook(lfHook)

	return nil
}

func setNull() {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	writer := bufio.NewWriter(src)
	log.SetOutput(writer)
}
